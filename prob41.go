package euler

func P41() int {
	a := []int{}
	return newL(0, a)
}

func newL(m int, a []int) int {
	for i := 1; i <= 7; i++ {
		t := copyArr(a)
		if !elem(i, t) {
			t = append(t, i)
			if len(t) == 7 || len(t) == 4 {
				r := digitsToNum(t)
				if len(t) == maximum(t) && r > m && isPrime(r) {
					m = r
				}
			}
			if len(t) < 7 {
				m = newL(m, t)
			}
		}
	}
	return m
}
