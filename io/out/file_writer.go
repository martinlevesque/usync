package out

import (
	// "fmt"
	"log"
	"os"
)

// FileWriter -
type FileWriter struct {
	Writer
	FilePtr *os.File
}

func (writer *FileWriter) Write(bytes []byte) (*Writing, string) {
	if writer.FilePtr == nil {
		file, err := os.Create(writer.URI)
		writer.FilePtr = file

		if err != nil {
			log.Fatal(err)
		}
	}

	count, err := writer.FilePtr.Write(bytes)
	writing := Writing{URI: writer.URI, BytesWritten: int64(count)}

	if err != nil {
		return &writing, err.Error()
	}

	return &writing, ""
}
