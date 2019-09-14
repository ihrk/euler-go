package problems

func P120() int {
	s := 0
	for i := 3; i <= 1000; i++ {
		s += i * (i - 2 + i%2)
	}
	return s
}
