package main

type intVector struct {
	x int
	y int
}

func (a intVector) Plus(b intVector) intVector {
	return intVector{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (a intVector) Minus(b intVector) intVector {
	return intVector{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a intVector) Multiply(b int) intVector {
	return intVector{
		x: a.x * b,
		y: a.y * b,
	}
}

func (a intVector) Divide(b int) intVector {
	return intVector{
		x: a.x / b,
		y: a.y / b,
	}
}
