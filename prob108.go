package euler

func P108() int {
	for i := 0; ; i++ {
		if (sqDivNum(i)+1)/2 > 1000 {
			return i
		}
	}
}

func sqDivNum(n int) int {
	s := 1
	if n == 0 {
		return s
	}
	sq := sqrt(n)
	if n%2 == 0 {
		t := 1
		for n%2 == 0 {
			t += 2
			n /= 2
		}
		s *= t
	}
	for i := 3; i <= sq; i += 2 {
		if n%i == 0 {
			t := 1
			for n%i == 0 {
				t += 2
				n /= i
			}
			s *= t
		}
	}
	if n != 1 {
		s += 2
	}
	return s
}
