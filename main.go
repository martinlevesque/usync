package main

import (
	// "fmt"
	"flag"
	"log"

	"github.com/martinlevesque/usync/io/in"
	"github.com/martinlevesque/usync/io/out"
	// "time"
)

const invalidInNetPort = -1

type commandArgs struct {
	stdin           bool
	stdout          bool
	input_type      string
	input_filepath  string
	output_filepath string

	in_net_port int
}

func (args *commandArgs) ParseArgs() {
	flag.BoolVar(&args.stdin, "stdin", false, "Stream content from stdin.")
	flag.StringVar(&args.input_filepath, "in-file", "", "Input filepath (file).")
	flag.IntVar(&args.in_net_port, "in-net-port", invalidInNetPort, "UDP in port")

	flag.BoolVar(&args.stdout, "stdout", false, "Stream content to stdout.")
	flag.StringVar(&args.output_filepath, "out-file", "", "Output filepath (file).")

	flag.Parse()

	// validate arguments

	// requires IN
	if len(args.input_filepath) == 0 && !args.stdin && args.in_net_port == invalidInNetPort {
		flag.Usage()
		log.Fatal("Invalid command line arguments, requires an input parameter")
	}

	// requires OUT
	if len(args.output_filepath) == 0 && !args.stdout {
		flag.Usage()
		log.Fatal("Invalid command line arguments, requires an output parameter")
	}
}

// Main function
func main() {
	// start := time.Now()

	args := &commandArgs{}

	args.ParseArgs()

	var reader in.IReader

	if args.stdin {
		reader = &in.StdinReader{Reader: in.Reader{URI: ""}}
	} else if args.in_net_port != invalidInNetPort {
		reader = &in.NetReader{Reader: in.Reader{URI: args.input_filepath}, Port: args.in_net_port}
	} else {
		reader = &in.FileReader{Reader: in.Reader{URI: args.input_filepath}, Index: 0}
	}

	var writer out.IWriter

	if args.stdout {
		writer = &out.StdoutWriter{Writer: out.Writer{URI: args.output_filepath}}
	} else {
		writer = &out.FileWriter{Writer: out.Writer{URI: args.output_filepath}}
	}

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
}
