package main

import (
	"fmt"
	"main/worker"
	"time"
)

func main() {
  // assuming that this is an asynchronous process and takes time
  defer worker.Wait()


  for i := 0; i< 10; i ++ {
	job := worker.Job{
    Action: PrintPayload,
    Payload: map[string]string{
      "time": time.Now().String(),
  }
  }
  job.Queue()
}
}

// code for action function

package main

import (
  "fmt"
  "time"
)

func PrintPayload(payload map[string]string){
  time.Sleep(time.Second * 3)
  fmt.Print(payload)
}


// code for worker

package worker

import "sync"

var wg sync.WaitGroup

type Job struct {
  id string
  Action func(map[string]string)
  Payload map[string]string
}

// to make the execution asynchrnous
func Wait (){
  wg.Wait()
}

func Queue()(job Job){
  wg.Add()
  go func
   defer wg.Done()
    job.Action(job.Payload)
  }()
  
}




