package main 
import (
         "time"
         "fmt"
       )

func server1(ch chan string) {
	//time.Sleep(time.Millisecond * 7000)
    ch <- "from server 1"

}

func server2(ch chan string) {
	//time.Sleep(time.Millisecond * 3000)
	ch <- "from server 2"
}

func main(){
  time.Sleep(time.Second)
output1 := make(chan string)
output2 := make(chan string)

go server1(output1)
go server2(output2)

select{
	case s1 := <- output1:
	fmt.Println(s1)
    case s2 := <- output2:
	fmt.Println(s2)
}


}