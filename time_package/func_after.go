package main
import "fmt"
import "time"
var ch chan int
func func_after(int){
fmt.Println("hii")
}
func main() {
select {
case m:=<-ch:
func_after(m)
case <-time.After(10 * time.Second):
fmt.Println("session timed out")
}}
