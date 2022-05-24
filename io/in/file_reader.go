
package in

import (
	"fmt"
)

type FileReader struct {
	Reader
}

func (r FileReader) Read() Reading {
    fmt.Println("file reading...")

    return Reading { Uri: r.Uri, Content: []byte {1,2} }
}
