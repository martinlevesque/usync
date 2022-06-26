package in

import (
	// "fmt"
	"log"
	"os"
)

type FileReader struct {
	Reader
	FilePtr   *os.File
	Index     int64
	ByteBlock []byte
}

func (reader *FileReader) Read() (*Reading, string) {

	if reader.FilePtr == nil {
		file, err := os.Open(reader.URI)
		reader.FilePtr = file
		reader.ByteBlock = make([]byte, StreamBlockSize)

		if err != nil {
			log.Fatal(err)
		}
	}

	reading := Reading{URI: reader.URI, Content: reader.ByteBlock}

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
