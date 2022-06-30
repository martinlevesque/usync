package out

import (
	// "fmt"
	"bufio"
	"os"
	// "io"
)

// StdoutWriter -
type StdoutWriter struct {
	Writer
}

func (writer *StdoutWriter) Write(bytes []byte) (*Writing, string) {

	bufWriter := bufio.NewWriter(os.Stdout)
	defer bufWriter.Flush()
	count, err := bufWriter.Write(bytes)

	writing := Writing{URI: writer.URI, BytesWritten: int64(count)}

	if err != nil {
		return &writing, err.Error()
	}

	return &writing, ""
}
