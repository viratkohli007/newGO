package main 
import ( 
    	"fmt"
    	
       )


func main(){

    
   ch := make(chan int)
   ch2 := make(chan int)

   fmt.Println("channel:", ch)
	go func () {
		i := 10
		ch <- i
		// fmt.Println(<-ch)
	    
	}()

	go func () {
		r := <-ch
		r2 := r+1;
      ch2 <- r2
      
	}()
      r2 := <-ch2 

      fmt.Println(r2)
	
}