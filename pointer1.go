package main
import  "fmt"
func main() {
a:=200
var b *int = &a
fmt.Println("address of a :",b)
fmt.Println("value of a :",a)
fmt.Println("address of pointer  b :",b)
fmt.Println("value of pointer b :", *b)
*b= *b+1
fmt.Println("vallue of b",*b)
//prints the new value using a and *b
fmt.Println("vallue of a",a)
}
