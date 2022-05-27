
package in

import (
	// "fmt"
	"os"
	"log"
)

const STREAM_BLOCK_SIZE int = 10

type FileReader struct {
	Reader
	FilePtr *os.File
	Index int64
}

func (reader *FileReader) Read() (*Reading, string) {
	file, err := os.Open(reader.Uri)
	reader.FilePtr = file

	if err != nil {
		log.Fatal(err)
	}

	reading := Reading { Uri: reader.Uri, Content: make([]byte, STREAM_BLOCK_SIZE) }

	count, err := file.ReadAt(reading.Content, reader.Index)

	if count == 0 {
		return &reading, "EOF"
	}

	if err != nil {
		return &reading, err.Error()
	}

	reader.Index += int64(count)

    return &reading, ""
}
