package out

// Writer base object
type Writer struct {
	URI string
}

// Writing - used to described the state of writes
type Writing struct {
	URI          string // NOLINT
	BytesWritten int64
}

// IWriter -
type IWriter interface {
	Write(bytes []byte) (*Writing, string)
}

func (writer *Writer) Write() (*Writing, string) {
	return &Writing{}, ""
}
