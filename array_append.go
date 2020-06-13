package main
import "fmt"
func main() {
a := [5] string {"one","two","three","four","five"}
fmt.Println("After creating A :",a)
slice_a := a[2:4]
fmt.Println("after slicing A :", slice_a)
b := [5] string {"1","2","3","4","5"}
fmt.Println("After creating B :",b)
slice_b := b[2:4]
fmt.Println("after slicing B :", slice_b)
//
slice_a = append(slice_a, slice_b...)
fmt.Println("after appending  :", slice_a)
}
