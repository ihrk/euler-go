package euler

func P37() int {
	s := 0
	for i, c := 1, 0; c < 11; i++ {
		if isTrunc(i) {
			s += i
			c++
		}
	}
	return s
}

func isTrunc(n int) bool {
	if n < 10 {
		return false
	}
	if !isPrime(n) {
		return false
	}
	lt := pow10(log(n, 10))
	for ; lt > 1; lt /= 10 {
		if !isPrime(n % lt) {
			return false
		}
	}
	for n /= 10; n > 0; n /= 10 {
		if !isPrime(n) {
			return false
		}
	}
	return true
}
