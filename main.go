package main

import (
	"fmt"
)

func main() {
	x:= soma(1,2,3)
	y:= multiplica(10,10)
	w:= subtrai(5, 10)
	z:= divide(20)
	fmt.Println(x, y, w, z)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func multiplica(x ...int) int {
	multiplica := 1
	for _, v := range x {
		multiplica *= v
	}
	return multiplica
}

func subtrai(x ...int) int {
	subtrai := 0
	for _, v := range x {
		subtrai -= v
	}
	return subtrai
}

func divide(x ...float32) float32 {
	var divide float32 = 1
	for _, v := range x {
		divide /= float32(v)
	}
	return divide
}
