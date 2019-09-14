package euler

func P46() int {
	for i := 35; ; i += 2 {
		if isPrime(i) {
			continue
		}
		b := false
		for j := 1; j*j*2 < i; j++ {
			if isPrime(i - 2*j*j) {
				b = true
				break
			}
		}
		if b {
			continue
		}
		return i
	}
}
