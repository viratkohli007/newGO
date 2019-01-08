package main 
import ( 
       "sync"
       "fmt"
        )

 var wg = sync.WaitGroup{}

func main(){
   ch := make(chan int)
   fmt.Println("hii")

   for i := 0; i < 10; i++ {
   	
        wg.Add(2)
   	    go func () {
   	 p := <- ch
   	 fmt.Println(p)
     wg.Done()
    }()

    go func () {
   	     var i = 10
   	     ch <- i
        wg.Done()
        }()
   }
 wg.Wait()
}

