
package in

type Reading struct {
    path string
    content []byte
}

type SyncReader interface {
    Read() Reading
}
