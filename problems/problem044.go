package problems

import "github.com/ihrk/euler-go/utils"

func P44() int {
	pents := []int{}
	s := 0
	b := true
	for i := 1; b; i++ {
		t := genPent(i)
		pents = append(pents, t)
		for j := len(pents) - 2; j >= 0; j-- {
			p := pents[j]
			if isPent(t-p) && isPent(t+p) {
				if s == 0 {
					s = t - p
					b = false
					break
				}
			}
		}
	}
	return s
}

func isSqr(n int) bool {
	t := utils.SqrtInt(n)
	if t*t == n {
		return true
	}
	return false
}

func isPent(n int) bool {
	t := n*24 + 1
	if isSqr(t) && (1+utils.SqrtInt(t))%6 == 0 {
		return true
	}
	return false
}

func genPent(n int) int {
	return n * (3*n - 1) / 2
}
