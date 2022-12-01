package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	arr[0] = 10
	fmt.Println(arr)

	for i:=0;i<len(arr);i++{
		fmt.Print(arr[i],"\n")
	}
}