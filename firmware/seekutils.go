package firmware

import "bufio"

func emptySeek(r *bufio.Reader, length int) {
	for i := 0; i < length; i++ {
		// just move forward the pointer
		r.ReadByte()
	}
}

func seekAndRead(r *bufio.Reader, length int, reversed bool) []byte {
	data := []byte{}

	for i := 0; i < length; i++ {
		d, _ := r.ReadByte()
		data = append(data, d)
	}

	if reversed {
		return reverse(data)
	}

	return data
}
