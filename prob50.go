package euler

func P50() int {
	const lim = 1000000
	primes := primeSieve(lim)
	lg := 0
	sum := 0

	for i, s := 0, 0; i < len(primes); i++ {
		s += primes[i]
		if s >= lim {
			break
		}
		lg++
		sum = s
	}
	if isPrime(sum) {
		return sum
	}
	for i, s := 0, sum; ; {
		s -= primes[i]
		s += primes[i+lg]
		if s > lim {
			sum -= primes[lg-1]
			lg--
			s = sum
			i = 0
			if isPrime(s) {
				break
			}
		} else if isPrime(s) {
			sum = s
			break
		} else {
			i++
		}
	}
	return sum
}
