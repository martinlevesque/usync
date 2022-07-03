package in

import (
	"fmt"
	"log"
	"net"
)

// NetReader -
type NetReader struct {
	Reader
	ByteBlock    []byte
	Port         int
	ConnListener net.Listener
	CurrentConn  net.Conn
	Initialized  bool
}

func (reader *NetReader) initializeServer() {
	log.Println("Will listen on port")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", reader.Port))
	reader.ConnListener = listener

	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	conn, errAccept := listener.Accept()
	reader.CurrentConn = conn

	if errAccept != nil {
		log.Fatal(err)
	}

	log.Println("Initialized server on port ", reader.Port)
}

func (reader *NetReader) Read() (*Reading, string) {

	if !reader.Initialized {
		reader.ByteBlock = make([]byte, StreamBlockSize)
		reader.initializeServer()
		reader.Initialized = true
	}

	// TODO make it accept connections

	reading := Reading{URI: "", Content: reader.ByteBlock}

	rlen, err := reader.CurrentConn.Read(reader.ByteBlock)
	reading.Length = int64(rlen)

	if rlen == 0 {
		reader.Initialized = false
		return &reading, "EOC"
	}

	if err != nil {
		reader.Initialized = false
		return &reading, "EOC"
	}

	return &reading, ""
}
