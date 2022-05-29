
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

	_, err := writer.FilePtr.Write(bytes)

	if err != nil {
		return nil, err.Error()
	}

	// fmt.Println("n read:")
	// fmt.Println(n)

	return nil, ""
}
