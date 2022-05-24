
package main

import (
	"fmt"
	"github.com/martinlevesque/usync/io/in"
)
  
// Main function
func main() {

	//in.Hello()
	var reader in.IReader
	reader = in.FileReader{ in.Reader { Uri: "/usr/bin/what" } }
	reading := reader.Read()
  
    fmt.Println("!... Hello World ...!")
    fmt.Println(reading)
}
