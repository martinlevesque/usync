
package in

type Reader struct {
	Uri string
}

type Reading struct {
    Uri string
    Content []byte
    Length int64
}

type IReader interface {
    Read() (*Reading, string)
}

func (b *Reader) Read() (*Reading, string) {
	return &Reading{}, ""
}
