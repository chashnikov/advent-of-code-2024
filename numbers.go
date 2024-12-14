package main

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

func Gcd(x, y int) int {
	x = Abs(x)
	y = Abs(y)
	for x > 0 && y > 0 {
		x, y = max(x, y), min(x, y)
		x, y = y, x%y
	}
	return x + y
}
