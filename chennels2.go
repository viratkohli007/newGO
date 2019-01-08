package main 
import ( 
    	"fmt"
    	"sync"
       )

var wg = sync.WaitGroup{}

func main(){

    
   ch := make(chan int)
    wg.Add(2)
	go func () {
		i := 10
		ch <- i
		fmt.Println(<-ch)
	    wg.Done()
	}()

	go func () {
		r := <-ch
      fmt.Println(r)
      ch <- 12
      wg.Done()
	}()
	wg.Wait()
}