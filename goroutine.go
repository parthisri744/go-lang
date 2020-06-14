package main
import "fmt"
func display() {
for i :=1;i<=5;i++ {
fmt.Println("hai everone")
}
}
func main() {
go display()
for i:= 1; i<=5;i++ {
fmt.Println("hii good morining")
}
}
