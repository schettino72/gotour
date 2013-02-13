package main

import "fmt"

func Cbrt(x complex128) complex128 {
	var result complex128 = 1.0
	for i := 0; i < 10; i++ {
		result -= (result*result*result - x) / (3 * result * result)
	}
	return result

}

func main() {
	got := Cbrt(2)
	fmt.Println(got, got*got*got)
}
