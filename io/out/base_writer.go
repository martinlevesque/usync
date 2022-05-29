
package out

type Writer struct {
	Uri string
}

type Writing struct {
    Uri string
    BytesWritten int64
}

type IWriter interface {
    Write(bytes []byte) (*Writing, string)
}

func (writer *Writer) Write() (*Writing, string) {
	return &Writing{}, ""
}
