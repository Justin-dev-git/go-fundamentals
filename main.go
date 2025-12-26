package main

import "fmt"

func main() {
	var Zahl []int
	i := 1101
	for i > 0 {
		Zahl = append(Zahl, i%2)
		i /= 2
	}
	fmt.Println(Zahl)
	s := 0.375
	var Zahl2 []int
	for s != 1 {
		Zahl2 = append(Zahl2, int(s*2))
		s *= 2
		if s > 1 {
			s -= 1
		}
	}
	fmt.Println(Zahl2)
}
