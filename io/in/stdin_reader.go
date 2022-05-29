
package in

import (
	// "fmt"
	"bufio"
	"os"
	"io"
)

type StdinReader struct {
	Reader
	ByteBlock []byte
	BlockInitialized bool
}

func (reader *StdinReader) Read() (*Reading, string) {

	if ! reader.BlockInitialized {
		reader.ByteBlock = make([]byte, STREAM_BLOCK_SIZE)
		reader.BlockInitialized = true
	}

	reading := Reading { Uri: "", Content: reader.ByteBlock }

	bufReader := bufio.NewReader(os.Stdin)
	count, err := bufReader.Read(reader.ByteBlock)
	reading.Length = int64(count)

	if count == 0 || err == io.EOF {
		return &reading, "EOF"
	}

	if err != nil {
		return &reading, err.Error()
	}

    return &reading, ""
}
