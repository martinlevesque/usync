
package main

import (
	"fmt"
	"github.com/martinlevesque/usync/io/in"
)
  
// Main function
func main() {

	//in.Hello()
	fReader := in.FileReader{}
	fReader.Read()
  
    fmt.Println("!... Hello World ...!")
}
