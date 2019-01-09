package main 
import "fmt"

func main() {
 
 ch1 := make(chan string, 2)
	

	
		 try(ch1)
	

}

func try(ch chan string) {
	ch <- "mayank"
    ch <- "darling"
    //ch <- "hdfh"



	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}