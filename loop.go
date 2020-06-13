package main
import "fmt"
func main() {
var i int
//for loop
for i=1;i<=10;i++ {
fmt.Println(i)
}

//if else
var x = 50
    if x < 10 {
        //Executes if x is less than 10
        fmt.Println("x is less than 10")
    } else {
        //Executes if x >= 10
        fmt.Println("x is greater than or equals 10")
    }
//nested if statement
var y int
y = 20
if y < 20 {
fmt.Println("y is less then 20")
}else if  y !=21 {
fmt.Println(" y is equal to 20 and greater then 50 ")
}else {
fmt.Println(" y is greater then 50")
}
//switch statement
  a,b := 3,1
    switch a+b {
    case 1:
        fmt.Println("Sum is 1")
    case 2:
        fmt.Println("Sum is 2")
    case 3:
        fmt.Println("Sum is 3")
    default:
        fmt.Println("Printing default")
    }
}

