package main
import "fmt"
func main() {
var number [3] string
number[0]="one"
number[1]="two"
number[2]="three"
fmt.Println(number)
fmt.Println(len(number))
fmt.Println(number[1])
directions:=[...] int {1,2,3,4,5}
fmt.Println(directions)
fmt.Println(len(directions))

//array modification
var a [5] string
a[0]="one"
a[1]="two"
a[2]="three"
a[3]="four"
a[4]="five"
fmt.Println("After creating Array :", a)
//array
 b := [3] string {"parthi","parthisri","parthiban"}
 fmt.Println("second array :" ,b)
 //
 var c [] string = a[2:4]
 fmt.Println("after slicing array :",c)
 c[1] = "parthisri"
 fmt.Println("slice after modification",c)
 fmt.Println("slice count :",len(c))
 }
 
