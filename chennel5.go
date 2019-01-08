package main 
import (
      	"fmt"
      	"sync"
      	)

var x = 0
func main() {
	
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}
	wg.Wait()
	fmt.Println("the value of x is: ", x)
}

func increment(wg *sync.WaitGroup, ch chan bool) {
	
	ch <- true
	x = x+1
	<-ch
	wg.Done()
}