package main
import (
         "fmt"
         "sync"
         "runtime"
        )

var counter = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main(){

	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))
 
    for i := 0; i < 10; i++ {

    	  wg.Add(2)
    	  m.RLock()
    	    go sayHello()
    	    m.Lock()
    	   go increment()
    	 

    	
    	  
    }
    wg.Wait()
}

func increment() {
      
	   counter++
	   m.Unlock()
	   wg.Done()
}

func sayHello(){
     
	fmt.Println("Hello ", counter)
	m.RUnlock()
	wg.Done()
}