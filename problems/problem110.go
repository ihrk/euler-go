package problems

import (
	"sort"

	"github.com/kavaliouIhar/euler-go/utils"
)

func P110() int {
	m := dioph(4000000)
	a := utils.PrimeGen(utils.Log(m, 3) + 1)
	a = optimizeDioph(m, a)
	return utils.Product(a)
}

func optimizeDioph(n int, a []int) []int {
	c := utils.CopyArr(a[:len(a)-1])
	m := a[len(a)-1]
	for i := 2; i < m; i++ {
		t := append(c, utils.PrimeFactors(i)...)
		sort.Ints(t)
		if n <= getDioph(utils.Group(t)) {
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
