package euler

import (
	"sort"
)

func P110() int {
	m := dioph(4000000)
	a := primeGen(log(m, 3) + 1)
	a = optimizeDioph(m, a)
	return product(a)
}

func optimizeDioph(n int, a []int) []int {
	c := copyArr(a[:len(a)-1])
	m := a[len(a)-1]
	for i := 2; i < m; i++ {
		t := append(c, primeFactors(i)...)
		sort.Ints(t)
		if n <= getDioph(group(t)) {
			t = optimizeDioph(n, t)
			return t
		}
	}
	return a
}

func getDioph(d [][]int) int {
	s := 1
	for _, v := range d {
		s *= len(v)*2 + 1
	}
	return s
}

func dioph(n int) int {
	return 2*n - 1
}
