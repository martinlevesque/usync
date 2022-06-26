package in

const StreamBlockSize int = 1000000

type Reader struct {
	URI string
}

type Reading struct {
	URI     string
	Content []byte
	Length  int64
	What    int64
}

type IReader interface {
	Read() (*Reading, string)
}

func (b *Reader) Read() (*Reading, string) {
	return &Reading{}, ""
}
