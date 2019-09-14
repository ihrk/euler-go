package euler

func P32() int {
	a := []int{}
	for i := 2; i < 100; i++ {
		for j := 100; ; j++ {
			t := i * j
			k := []int{}
			k = append(k, numToDigits(i)...)
			k = append(k, numToDigits(j)...)
			k = append(k, numToDigits(t)...)
			if len(k) > 9 {
				break
			}
			if len(k) == 9 && noRepeat(k) && !elem(0, k) {
				if !elem(t, a) {
					a = append(a, t)
				}
			}

		}
	}
	return sum(a)
}

func noRepeat(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}
