package problems

import (
	"math/big"

	"github.com/kavaliouIhar/euler-go/utils"
)

func P66() int {
	s := 0
	max := big.NewInt(0)
	for i := 2; i <= 1000; i++ {
		if t := diophSolveBig(i); t.Cmp(max) == 1 {
			max = t
			s = i
		}
	}
	return s
}

func diophSolveBig(n int) *big.Int {
	a := sqrtToFracSeq(n)
	if len(a) < 2 {
		return big.NewInt(0)
	}
	x := big.NewInt(1)
	y := big.NewInt(int64(a[len(a)-1]))
	for i := len(a) - 2; i >= 0; i-- {
		t := new(big.Int)
		t.Mul(big.NewInt(int64(a[i])), y)
		x.Add(x, t)
		x, y = y, x
	}
	if len(a)%2 != 0 {
		y.Mul(y, y)
		y.Mul(y, big.NewInt(2))
		return y.Add(y, big.NewInt(1))
	}
	return y
}

func sqrtToFracSeq(n int) []int {
	base := utils.SqrtInt(n)
	if isSqr(n) {
		return []int{base}
	}
	if isSqr(n - 1) {
		return []int{base, 2 * base}
	}
	a := []int{base}
	y := n - base*base
	x := base
	for {
		t1 := (x + base) / y
		t2 := (x + base) % y
		if t1 == base*2 {
			break
		}
		a = append(a, t1)
		x = base - t2
		y = (n - x*x) / y
	}
	return a
}
