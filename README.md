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
-------------------------------


package main

import (
"fmt"
"math"
)

func main() {
fmt.Println(maximum()) 
}

func maximum() int {
	i := 1;
	for  {
		x := math.Pow(0.9, float64(i+1)) - 1/float64(i+1)
		y := math.Pow(0.9, float64(i)) - 1/float64(i)
		if !((x - y) >= 0) {
			return i
		}
		i++
	}
}

