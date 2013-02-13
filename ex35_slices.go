package main

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dx; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			ret[i][j] = uint8(i ^ j + (i+j)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}