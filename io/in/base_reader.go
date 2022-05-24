
package in

type Reader struct {
	Uri string
}

type Reading struct {
    Uri string
    Content []byte
}

type IReader interface {
    Read() Reading
}

func (b *Reader) Read() Reading {
	return Reading{}
}
