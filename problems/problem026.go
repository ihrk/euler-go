package problems

import "math/big"

func P26() int {
	max := 0
	s := 0
	for i := 2; i < 1000; i++ {
		t := cycleLen(i)
		if t > max {
			max = t
			s = i
		}
	}
	return s
}

func cycleLen(n int) int {
	if n < 2 {
		return 0
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%5 == 0 {
		n /= 5
	}
	if n == 1 {
		return 0
	}
	r := 1
	i := big.NewInt(10)
	b := big.NewInt(int64(n))
	one := big.NewInt(1)
	ten := big.NewInt(10)
	for {
		t := new(big.Int)
		if one.Cmp(t.Mod(i, b)) == 0 {
			return r
		}
		i.Mul(i, ten)
		r++
	}
}
