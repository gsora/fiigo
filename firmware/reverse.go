package firmware

func reverse(orig []byte) []byte {
	rt := make([]byte, len(orig))

	for i := (len(orig) - 1); i >= 0; i-- {
		rt[(len(rt)-1)-i] = orig[i]
	}

	return rt
}
