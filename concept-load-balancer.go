package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// interface is like a contract in golang
// a function on the interface has access to all the functions
type Server interface{
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// simple server with proxy and port address
// proxy forwards the request to another server
type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

// function to create a new server
// returns a simple server
func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		fmt.Println(err)
	}

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

// load balancer structs that does the forwading
type LoadBalancer struct {
	port string
	roundRobinCount  int
	servers []Server
}

// function to create a new load balancer
func newLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: port,
		roundRobinCount: 0,
		servers: servers,

	}
}

// Methods on the struct simple server

// returns the address
func ( s *simpleServer) Address() string {return s.addr}

// checks if the server is alive or not
func ( s *simpleServer) IsAlive() bool {return true}

// wrapper arround the proxy function
func ( s *simpleServer) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r)
}


// method on the load balancer to increment and give the next server
func (lb *LoadBalancer) getNextAvailableServer() Server{
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive(){
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// method in the load balancer to forward it to the next server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request){
	targetServer := lb.getNextAvailableServer()
	fmt.Println("forwading request to", targetServer.Address())
	targetServer.Serve(rw, r)
}

func main(){
	// since simple sever satisfies the interface
	servers := []Server{
		newSimpleServer("https://wwww.facebook.com"),
		newSimpleServer("http//wwww.bing.com"),
		newSimpleServer("http://wwww.duckduckgo.com"),
	}

	lb := newLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request){
		lb.serveProxy(rw, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Println("serving requests at localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}

