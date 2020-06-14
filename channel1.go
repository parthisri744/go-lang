package main
import "fmt"
import "time"
func add(ch chan int) {
fmt.Println("send data") 
for i:=1;i<=5;i++ {
ch <-i
}
}
func fetch(ch chan int) {
fmt.Println("read data")
for {
x, flag :=<-ch
if flag==true {
fmt.Println(x)
}else {
fmt.Println("channel is empty")
}
}
}
func main() {
ch :=make(chan int)
go add(ch)
go fetch(ch)
time.Sleep(5*time.Second)
fmt.Println("inside main")
}
