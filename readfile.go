package main
import "fmt"
import "io/ioutil"
func main() {
data, err :=ioutil.ReadFile("About.txt")
if err!=nil {
fmt.Println("File reading error",err)
return
}
fmt.Println("content of the file",string(data))
}
