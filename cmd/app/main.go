package main

import "fmt"

func main(){
	/*
	for i:=0; i<5;i++{
		fmt.Println("Zähler: ", i)
	}
	i :=0
	for i<3{
		fmt.Println("Zähler: ",i+5)
		i++
	}
	for i:=0;i<10;i++{
		if i==4{
			break
		}
		fmt.Println("Zähler: ", i+8)
	}*/
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 12
	numbers[2] = 1
	numbers2 :=[3]int{1,2,3}

	fmt.Print(numbers)
	fmt.Println(numbers2)
}
