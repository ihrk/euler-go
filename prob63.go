package euler

func P63() int {
	s := 0
	for i := float64(1); i < 10; i++ {
		p := i / 10
		t := p
		for t >= 0.1 {
			s++
			t *= p
		}
	}
	return s
}
