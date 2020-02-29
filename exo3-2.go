package main

import ( "fmt"
		     "math" )


func maximum() int {
	i := 1; 
	for {     x := math.Pow(0.9, float64(i+1)) - 1/float64(i+1) 
        	  y := math.Pow(0.9, float64(i)) - 1/float64(i) 
    	    if !((x - y) >= 0) {
            	  return i 
        }
        	    i++ 
    } 
}

func main() { fmt.Println(maximum()) }
