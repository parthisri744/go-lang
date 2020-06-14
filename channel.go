/*Channels
channel_variable := make(chan datatype)
example:   ch := make(chan int)
You can send data to a channel using the syntax

channel_variable <- variable_name
Example

    ch <- x 
You can receive data from a channel using the syntax

    variable_name := <- channel_variable
Example

   y := <- ch
    */
package main
import "fmt"
import "time"
func display(ch chan int) {
time.Sleep(10*time.Second)
fmt.Println("iam parthiban")
ch <-1234
}
func  main() {
ch :=make(chan int)
go display(ch)
x := <-ch
fmt.Println("good moring")
fmt.Println("Printing x in main() after taking from channel:",x)
}

