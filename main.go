
package main

import (
	// "fmt"
	"github.com/martinlevesque/usync/io/in"
	"github.com/martinlevesque/usync/io/out"
	"log"
	"flag"
	// "time"
)

type CommandArgs struct {
    input_type string
    input_filepath string
    output_filepath string
}

func (args *CommandArgs) ParseArgs() {
	flag.StringVar(&args.input_filepath, "in-file", "", "Input filepath (file).")
	flag.StringVar(&args.output_filepath, "out-file", "", "Output filepath (file).")

	flag.Parse()

	// validate arguments

	if len(args.input_filepath) == 0 || len(args.output_filepath) == 0 {
		flag.Usage()
		log.Fatal("Invalid command line arguments")
	}
}
  
// Main function
func main() {
	// start := time.Now()

	args := &CommandArgs{}

	args.ParseArgs()

	var reader in.IReader
	reader = &in.FileReader{ Reader: in.Reader { Uri: args.input_filepath }, Index: 0 }

	var writer out.IWriter
	writer = &out.FileWriter{ Writer: out.Writer { Uri: args.output_filepath } }

	for {
		reading, state := reader.Read()

		if state == "EOF" {
			writer.Write(reading.Content[0:reading.Length])
			break
		} else if state != "" {
			log.Fatal(state)
		}

	    
	    writer.Write(reading.Content[0:reading.Length])
	}

	// elapsed := time.Since(start)
    // log.Printf("took %s", elapsed)
}
