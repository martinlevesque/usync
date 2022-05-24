
package in

import (
	"fmt"
)

type FileReader struct {
	dummy string
}

func (r FileReader) Read() {
    fmt.Println("file reading...")
}
