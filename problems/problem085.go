package problems

import "github.com/ihrk/euler-go/utils"

func P85() int {
	const lim = 2000000
	sq := utils.SqrtInt(utils.SqrtInt(lim * 4))

	d := lim
	s := 0
	for i := 1; i <= sq; i++ {
		j := utils.SqrtInt((lim * 4) / (i * (i + 1)))
		if r := rectNum(i, j); utils.Abs(lim-r) < d {
			s = i * j
			d = utils.Abs(lim - r)
		}
	}
	return s
}

func rectNum(a, b int) int {
	return a * (a + 1) * b * (b + 1) / 4
}
