package main 
import "fmt"
 //import "time"

func main() {

ch := make(chan int, 2)

go write(ch)

 //time.Sleep(time.Millisecond * 300)
for v := range ch {
     fmt.Println("read value",v, " from ch")
    // time.Sleep(time.Millisecond *3)
}
	
}

func write(ch chan int) {
	 
for i := 0; i < 5; i++ {
	ch <- i
	fmt.Println("successfully wrote", i, "to the channel")
}
close(ch)

}