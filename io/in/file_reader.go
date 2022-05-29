
package in

import (
	// "fmt"
	"os"
	"log"
)

const STREAM_BLOCK_SIZE int = 1000000

type FileReader struct {
	Reader
	FilePtr *os.File
	Index int64
	ByteBlock []byte
}

func (reader *FileReader) Read() (*Reading, string) {

	if reader.FilePtr == nil {
		file, err := os.Open(reader.Uri)
		reader.FilePtr = file
		reader.ByteBlock = make([]byte, STREAM_BLOCK_SIZE)

		if err != nil {
			log.Fatal(err)
		}
	}

	reading := Reading { Uri: reader.Uri, Content: reader.ByteBlock }

	count, err := reader.FilePtr.ReadAt(reading.Content, reader.Index)
	reading.Length = int64(count)

	if count == 0 {
		return &reading, "EOF"
	}

	if err != nil {
		return &reading, err.Error()
	}

	reader.Index += int64(count)

    return &reading, ""
}
