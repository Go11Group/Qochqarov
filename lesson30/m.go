package main

import (
	"fmt"
	"math"
)

func main() {
	k := 0
    var	str rune
	s := "lmdoewm"
	for i, v := range s {
		if i >= 1 && i < len(s)-1 {
			k +=int( math.Abs(float64(v - str)))
		}
		fmt.Println(v)
		str = v
	}
	fmt.Println(k)

}
