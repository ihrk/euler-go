package problems

func P72() int {
	const lim = 1e6
	a := make([]int, lim+1)
	for i := 2; i <= lim; i++ {
		if a[i] == 0 {
			for j := i * i; j <= lim; j += i {
				if a[j] == 0 {
					a[j] = i
				}
			}
		}
	}
	s := 0
	for i := 2; i <= lim; i++ {
		s += eulerPhi(i, a)
	}
	return s
}

func eulerPhi(n int, a []int) int {
	if n <= 1 {
		return 1
	}
	f := a[n]
	if f == 0 {
		return n - 1
	}
	n /= f
	fp := 1
	for n%f == 0 {
		fp *= f
		n /= f
	}
	return eulerPhi(n, a) * (f - 1) * fp
}
