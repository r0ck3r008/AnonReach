package utils

func Getlvl(s1 *string, s2 *string) int {
	var count int = 0
	var mlen int
	if len(*s1) < len(*s2) {
		mlen = len(*s1)
	} else {
		mlen = len(*s2)
	}

	for ; count < mlen || (*s1)[count] != (*s2)[count]; count++ {
	}

	return count
}
