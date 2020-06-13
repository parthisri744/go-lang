package main
import "fmt"
func number(num1 int,num2 int)(int, int, int) {
sum:=num1+num2
diff:= num1-num2
mul:=num1*num2
return sum,diff,mul
}
func main() {
x,y := 10,5
a,b,c:=number(x,y)
fmt.Println("sum of x and y is :" ,a)
fmt.Println("diff between two numbers is :",b)
fmt.Println("multiplication of two no is :",c)
}
