package main 
import "fmt"

func main() {

	ch := make(chan int, 5)
	done := make(chan bool)

	  go func () {
	  	
         for {j, more:=  <-ch
         if more{
         	fmt.Println("recieved chennel ", j)
         	}else{
            fmt.Println("recieved all jobs")
            done <- true
            return
         	}
         	
         }

	  }()
     



	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Println("send chennel",i)
	}
	 close(ch)
	fmt.Println("sent all jobs")
	<-done
}