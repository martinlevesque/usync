
package main

import (
	"fmt"
	"github.com/martinlevesque/usync/io/in"
	"log"
)
  
// Main function
func main() {

	//in.Hello()
	var reader in.IReader
	reader = &in.FileReader{ Reader: in.Reader { Uri: "/home/martin/works/usync/README.md" }, Index: 0 }
	reading, state := &in.Reading{}, ""

	for {
		reading, state = reader.Read()

		if state == "EOF" {
			fmt.Print(string(reading.Content))
			break
		} else if state != "" {
			log.Fatal(state)
		}

		//fmt.Println("reading: ")
	    //fmt.Println(reading)
	    fmt.Print(string(reading.Content))
	}

}
