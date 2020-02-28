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
---------------------------
package main

import (
"fmt"
)

func main() {
fmt.Println(control(15, 1, 3, 0)) 
}

func control(age, permis, assurance, accident int) string {
	switch {
		case age < 25 && permis < 2:
			if accident == 0 {
				if assurance >= 5 {
					return "Orange"
				}
				return "Rouge"
			} else {
				return "Refusé"
			}
		case (age < 25 && permis > 2) || (age > 25 && permis < 2):
		
			if accident == 0 {
				if assurance >= 5 {
					return "Vert"
				}
				return "Orange"
			} else if accident == 1 {
				if assurance >= 5 {
					return "Orange"
				}
				return "Rouge"
			}
			return "Refusé"
		case (age > 25 && permis > 2):
			switch accident {
				case 0:
					if assurance >= 5 {
						return "Bleu"
					}
					return "Vert"
				case 1:
					if assurance >= 5 {
						return "Vert"
					}
					return "Orange"
				case 2:
					if assurance >= 5 {
						return "Orange"
					}
					return "Rouge"
				default:
					return "Refusé"
			}
		
		
		
	}
	return ""
}

