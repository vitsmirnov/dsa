package main

type BIT2D struct{ sums [][]int }

func MakeBIT2D(height, width int) *BIT2D {
	sums := make([][]int, height)
	for y := range sums {
		sums[y] = make([]int, width)
	}
	return &BIT2D{sums: sums}
}

func (bit *BIT2D) Sum(top, left, bottom, right int) int {
	sum := bit.sum(bottom, right)
	if top > 0 {
		sum -= bit.sum(top-1, right)
	}
	if left > 0 {
		sum -= bit.sum(bottom, left-1)
	}
	if top > 0 && left > 0 {
		sum += bit.sum(top-1, left-1)
	}
	return sum
}

func (bit *BIT2D) sum(bottom, right int) int {
	sum := 0
	for y := bottom; y >= 0; y = (y & (y + 1)) - 1 {
		for x := right; x >= 0; x = (x & (x + 1)) - 1 {
			sum += bit.sums[y][x]
		}
	}
	return sum
}

func (bit *BIT2D) Update(x, y int, delta int) {
	h, w := len(bit.sums), len(bit.sums[0])
	for ; y < h; y |= y + 1 {
		for _x := x; _x < w; _x |= _x + 1 {
			bit.sums[y][_x] += delta
		}
	}
}
