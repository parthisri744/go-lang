package main
import "fmt"
type emp struct {
name string
age int
address string
}
func display(e emp) {
fmt.Println("Employee details :" )
fmt.Println("employee name",e.name)
fmt.Println("employee age",e.age)
fmt.Println("employee address",e.address)
}
func main() {
var empdata1 emp
empdata1.name = "parthiban"
empdata1.age = 21
empdata1.address = "thanigai amman kovil st anupuram"
 
empdata2 :=emp{"parthisri", 21, "kalpakkam tamilnadu"}
defer   display(empdata1)
display(empdata2)
}
