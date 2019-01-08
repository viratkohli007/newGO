package main
import ("fmt"
        "sync"
)

var wg = sync.WaitGroup{}

func main(){
	
	msg := "hiiii"
	wg.Add(1)

	go func (msg string){

	 fmt.Println(msg)
	 wg.Done()
	}(msg)
     msg = "hello"
	
    wg.Wait()
}