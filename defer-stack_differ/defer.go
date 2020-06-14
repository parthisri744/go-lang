package main
import "fmt"
func sample() {
fmt.Println("inside the sample")
}
func main() {
defer sample()
fmt.Println("inside the main")
}
