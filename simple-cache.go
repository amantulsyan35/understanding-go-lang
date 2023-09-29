package main

// global cache object on which operations will be performed
  type Cache struct {
    // key value pair of the cache key - result
    // *result instead of result because
    // https://chat.openai.com/share/ad8292f8-9b9f-47b2-ba0b-3b782ddbe7d0
    cache map[string]*result
  }

  type result struct {
    err error
    // can be anything instead of byte depending on the function
    // byte - uint8
    values []byte
  }

  type Func func() ([]byte, error)

  // we need two methods 
  // GET AND CREATE

func NewCache() *Cache {
  //make is an inbuilt method used to initialize a map
  return map[string]{cache: make(map[string]*result)}
}

// (c * Cache) means it is a funciton on the struct 
// In go this is the way to define method on a particular struct

// Get() is the signature with parameters
// (*Cache, error) is the return type of the function
func (c *Cache) Get(key string, f Func) (*Cache, error) {
    result, ok := c.cache[key]
    if !ok {
      // case of a cache miss 
      // store the value in the cache object
      //initializing result
      res := &result
      res.values, res.error := f()
      c.cache[key] = res    
    }
  return result, nil
}

func main() {
// implement the testing fn
}
