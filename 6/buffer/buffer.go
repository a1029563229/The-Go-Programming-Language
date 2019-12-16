package buffer

type Buffer struct {
	buf     []byte
	initial [64]byte
}

func init() {
	var b Buffer
	b.buf = []byte{10}
}
