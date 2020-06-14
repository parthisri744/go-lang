package main
import "fmt"
import "os"
func fileopen(name string) {
f, err := os.Open(name)
if err!=nil {
fmt.Println(err)
return
}else {
fmt.Println("file opened", f.Name())
}
}
func main() {
fileopen("About.txt")
}
