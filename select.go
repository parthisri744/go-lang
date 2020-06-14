package main
import "fmt"
//import	"time"
func data1(ch chan string) {
//time.Sleep(3*time.Second)
ch<-"fetch data1"
}
func data2(ch chan string) {
//time.Sleep(5*time.Second)
ch<-"fetch data2"
}
func main() {
ch1 :=make(chan string)
ch2 :=make(chan string)
go data1(ch1)
go data2(ch2)
select {
case x := <- ch1:
fmt.Println(x)
case y := <- ch2:
fmt.Println(y)
}
}
