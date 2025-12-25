package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4}
	numbers = append(numbers, 5)
	sum :=0
	for i, value := range numbers{
		fmt.Println("Index:", i, "Value", value)
		sum+=value	
	}
	fmt.Println(sum)
}
