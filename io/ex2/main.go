package ex2

import "io"

type Reader interface {
	Read(p []byte) (b int, err error)
}

func ReadForm(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
