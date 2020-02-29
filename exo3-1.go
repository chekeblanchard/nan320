package main

import "fmt"

func somme(a float64) float64{
	var k , h float64 = 1 , 0

	for{
			h = h + 1/k
		 
	if  h > a {
		return k
		}
		k++ 
	}

}


func main() {
	T := 0.0

	fmt.Print("Entrez un nombre réel,:")
	fmt.Scan(&T)
	fmt.Println(somme(T))
}
