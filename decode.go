package base256

import (
	"bufio"
	"io"
)

// DecodeString returns the bytes represented by the base256 string s.
// Non-base256 runes are skipped silently.
func DecodeString(s string) []byte {
	src := []rune(s)
	dst := []byte{}

	for i := 0; i < len(src); i++ {
		decoded, ok := dectab[src[i]]
		if !ok {
			continue
		}
		dst = append(dst, decoded)
	}

	return dst
}

type decoder struct {
	reader *bufio.Reader
	err    error
}

// NewDecoder constructs a new base256 stream decoder. Data read from the returned
// reader is base256 decoded from r. Non-base256 runes are skipped silently.
func NewDecoder(r io.Reader) io.Reader {
	return &decoder{
		reader: bufio.NewReader(r),
	}
}

func (d *decoder) Read(c []byte) (int, error) {
	if d.err != nil {
		return 0, d.err
	}

	// the least count of runes to read is the length of c
	var i int
	for i = 0; i < len(c); {
		r, _, err := d.reader.ReadRune()
		if err != nil {
			d.err = err
			return i, err
		}
		decoded, ok := dectab[r]
		if !ok {
			continue
		}
		c[i] = decoded
		i++
	}

	return i, nil
}
