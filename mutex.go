package main
import ("fmt"
	"time"
	"strconv"
	//"sync"
	"math/rand")
//declare a mutex instance
//var mu sync.Mutex
var count =0
func process(n int) {
for i:=0;i<=5;i++ {
time.Sleep(time.Duration(rand.Int31n(2))*time.Second)
//mu.Lock()
temp :=count
temp++
time.Sleep(time.Duration(rand.Int31n(2))*time.Second)
count = temp
//mu.Unlock()
}
fmt.Println("count after i="+strconv.Itoa(n)+ "count :",strconv.Itoa(count))
}
func main()  {
for i:=1;i<4;i++ {
go process(i)
}
time.Sleep(25*time.Second)
fmt.Println("final count :",count)
}
