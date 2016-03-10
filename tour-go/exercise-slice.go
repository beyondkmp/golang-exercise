package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		r[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			r[i][j] = uint8(i ^ j)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
