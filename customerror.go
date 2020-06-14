package main
import "fmt"
import "os"
//importing error package
import "errors"
//function accepts a filename and tries to open it.
func fileopen(name string) (string, error) {
f, er :=os.Open(name)
 if er != nil {
        //created a new error object and returns it  
        return "", errors.New("Custom error message: File name is wrong")
    }else{
    	return f.Name(),nil
    }
}
func main() {
filename,err := fileopen("Abot.txt")
if err!=nil {
fmt.Println(err)
}else {
fmt.Println("file is opened", filename)
}
}

