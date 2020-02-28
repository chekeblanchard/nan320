package main

import (
	"fmt"
)

func main() {
	fmt.Println(somme(3))
}

func somme(a  float32) float32 {
	var k, h float32 = 1, 0
	for {
		h = h + 1/k
		if h > a {
			return k
		}
		k++
	}
	
}

