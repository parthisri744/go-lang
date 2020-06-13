package main
import (  
    "fmt"
    "packages/addition"
)

func main() {  
	x,y := 15,10
	//the package will have function Do_add()
sum := addition.Do_add(x,y)
fmt.Println("Sum",sum) 
}
