package main 
import "fmt"
//import "time"
//import "sync"

//var wg = sync.WaitGroup{}

func main() {
	
	ch := make(chan string)
	ch1 := make(chan string)
	//wg.Add(2)
	 go sender(ch)
	 go reciever(ch, ch1)
	// wg.Wait()
	 fmt.Println(<-ch1)
	 
}

func reciever(ch <- chan string, ch1 chan <- string ) {
	
	p := <-ch
	ch1 <- p
	
}

func sender(ch chan <- string){

	ch <- "some text"
	//wg.Done()
}