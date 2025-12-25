package main

import (
	"fmt"
	"errors"
)

func divide(a int, b int)(int, error){
	if b==0{
		return 0, errors.New("divison by 0")
	}
	return a/b, nil
	
}
func main(){
	result, err :=divide(1,1)
	if err!=nil{
	fmt.Println("Fehler:",err)
	return
	}
	fmt.Println("Ergebnis", result)
	
}
