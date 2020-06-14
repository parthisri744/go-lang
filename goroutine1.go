package main
import ("fmt"
	"time")
func display() {
for i:=1;i<=5;i++ {
time.Sleep( 1 * time.Second)
fmt.Println("hi ")
}
}
func main() {
go display()
for i:=1;i<=5;i++ {
time.Sleep(2*time.Second)
fmt.Println("good moring")
}
}

