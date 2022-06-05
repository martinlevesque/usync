
package out

import (
	// "fmt"
	"os"
	"log"
)

const STREAM_BLOCK_SIZE int = 1000000

type FileWriter struct {
	Writer
	FilePtr *os.File
}

func (writer *FileWriter) Write(bytes []byte) (*Writing, string) {
	if writer.FilePtr == nil {
		file, err := os.Create(writer.Uri)
		writer.FilePtr = file

		if err != nil {
			log.Fatal(err)
		}
	}

	count, err := writer.FilePtr.Write(bytes)
	writing := Writing { Uri: writer.Uri, BytesWritten: int64(count) }

	if err != nil {
		return &writing, err.Error()
	}

	return &writing, ""
}
