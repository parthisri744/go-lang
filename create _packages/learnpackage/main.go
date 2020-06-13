package main

import (  
    "fmt"
    "learnpackage/simpleinterest"
)

func main() {  
    fmt.Println("Simple interest calculation")
    p := 200.0
    r := 100.0
    t := 1.0
    si := simpleinterest.Calculate(p, r, t)
    fmt.Println("Simple interest is", si)
}
