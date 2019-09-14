package euler

func P85() int {
	const lim = 2000000
	sq := sqrt(sqrt(lim * 4))

	d := lim
	s := 0
	for i := 1; i <= sq; i++ {
		j := sqrt((lim * 4) / (i * (i + 1)))
		if r := rectNum(i, j); abs(lim-r) < d {
			s = i * j
			d = abs(lim - r)
		}
	}
	return s
}

func rectNum(a, b int) int {
	return a * (a + 1) * b * (b + 1) / 4
}
