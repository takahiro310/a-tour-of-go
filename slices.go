package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := range img {
		img[y] = make([]uint8, dx)
		for x := range img[y] {
			img[y][x] = uint8(y * x)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}
