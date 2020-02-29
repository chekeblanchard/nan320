package main

import (
	"fmt"
)

func control(age, permis, accident , assurance int) string {
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
	return "Refusé"

}

func main() {
	Age,Permis,Assurance,Accident := 0,0,0,0
	
	fmt.Print("Entrez l'âge:")
	fmt.Scan(&Age)
	fmt.Print("Entrez le nombre d'années de permis: ")
	fmt.Scan(&Permis)
	fmt.Print( "Entrez le nombre d'accidents: ")
	fmt.Scan(&Accident)
	fmt.Print( "Entrez le nombre d'années d'assurance: ")
	fmt.Scan(&Assurance)
	fmt.Println(control(Age,Permis,Accident,Assurance))
}
