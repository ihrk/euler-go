package euler

func P49() int {
	t := []int{}
	for i, condition := 1000, true; condition; i++ {
		if i == 1487 {
			continue
		}
		for j := 9; j < (10000-i)/2; j += 9 {
			a := i
			b := a + j
			c := b + j
			if isPrime(a) && isPrime(b) && isPrime(c) {
				as, bs, cs := numToDigits(a), numToDigits(b), numToDigits(c)
				if isPermutation(as, bs) && isPermutation(bs, cs) {
					t = append(as, bs...)
					t = append(t, cs...)
					condition = false
					break
				}

			}
		}
	}
	return digitsToNum(t)
}

func isPermutation(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i := range a1 {
		if elemCount(a1[i], a1) != elemCount(a1[i], a2) {
			return false
		}
	}
	return true
}

func elemCount(n int, a []int) int {
	s := 0
	for _, v := range a {
		if v == n {
			s++
		}
	}
	return s
}
