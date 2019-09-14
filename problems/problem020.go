package problems

import (
	"math/big"
)

func P20() int {
	fac := big.NewInt(1)
	for i := 2; i <= 100; i++ {
		t := big.NewInt(int64(i))
		fac.Mul(fac, t)
	}
	str := []rune(fac.String())
	s := 0
	for _, e := range str {
		s += int(e - '0')
	}
	return s
}
