package main

import (
	"flag"
	"log"

	"github.com/martinlevesque/usync/io/in"
	"github.com/martinlevesque/usync/io/out"
	// "time"
)

const invalidInNetPort = -1

type commandArgs struct {
	stdin          bool
	stdout         bool
	inputType      string
	inputFilepath  string
	outputFilepath string

	inNetPort int
}

func (args *commandArgs) ParseArgs() {
	flag.BoolVar(&args.stdin, "stdin", false, "Stream content from stdin.")
	flag.StringVar(&args.inputFilepath, "in-file", "", "Input filepath (file).")
	flag.IntVar(&args.inNetPort, "in-net-port", invalidInNetPort, "UDP in port")

	flag.BoolVar(&args.stdout, "stdout", false, "Stream content to stdout.")
	flag.StringVar(&args.outputFilepath, "out-file", "", "Output filepath (file).")

	flag.Parse()

	// validate arguments

	// requires IN
	if len(args.inputFilepath) == 0 && !args.stdin && args.inNetPort == invalidInNetPort {
		flag.Usage()
		log.Fatal("Invalid command line arguments, requires an input parameter")
	}

	// requires OUT
	if len(args.outputFilepath) == 0 && !args.stdout {
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
	readerType := ""

	if args.stdin {
		reader = &in.StdinReader{Reader: in.Reader{URI: ""}}
	} else if args.inNetPort != invalidInNetPort {
		reader = &in.NetReader{Reader: in.Reader{URI: args.inputFilepath}, Port: args.inNetPort}
		readerType = "net"
	} else {
		reader = &in.FileReader{Reader: in.Reader{URI: args.inputFilepath}, Index: 0}
	}

	var writer out.IWriter

	if args.stdout {
		writer = &out.StdoutWriter{Writer: out.Writer{URI: args.outputFilepath}}
	} else {
		writer = &out.FileWriter{Writer: out.Writer{URI: args.outputFilepath}}
	}

	for {
		reading, state := reader.Read()

		if state == "EOF" && readerType != "net" {
			writer.Write(reading.Content[0:reading.Length])
			break
		} else if state == "EOC" && readerType == "net" {
		} else if state != "" {
			log.Fatal(state)
		}

		writer.Write(reading.Content[0:reading.Length])
	}
}
