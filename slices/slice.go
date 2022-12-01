package main

import "fmt"

func main() {
	slice :=[]int{1,2,3}
	// fmt.Print(slice)
	slice[0] = 10
	// fmt.Print(slice)
	slice1:=append(slice,100)
	fmt.Print(slice1)
	slice2:=append(slice1,100,111,3,4,4,4)
	fmt.Print(slice2)

}