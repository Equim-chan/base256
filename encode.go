package base256

import (
	"bufio"
	"io"
)

// EncodeToString returns the encoded base256 string of src.
func EncodeToString(src []byte) string {
	dst := []rune{}
	table := int8(0)

	for i := 0; i < len(src); i++ {
		dst = append(dst, enctab[table][src[i]])
		table = 1 - table
	}

	return string(dst)
}

type encoder struct {
	writer *bufio.Writer
	table  int8
	err    error
}

// NewEncoder returns a new base256 stream encoder. Data written to the returned
// writer will be encoded using base256 and then written to w.
// When finished writing, the caller must Close the returned encoder to flush
// any partially written blocks.
func NewEncoder(w io.Writer) io.WriteCloser {
	return &encoder{
		writer: bufio.NewWriter(w),
		table:  0,
	}
}

func (e *encoder) Write(c []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var i int
	for i = 0; i < len(c); i++ {
		if _, err := e.writer.WriteRune(enctab[e.table][c[i]]); err != nil {
			e.err = err
			return i, err
		}
		e.table = 1 - e.table
	}

	return i, nil
}

func (e *encoder) Close() error {
	if e.err != nil {
		return e.err
	}
	e.err = e.writer.Flush()

	return e.err
}
