package main

import "fmt"

func f1(a int){
	fmt.Print("Hello f1 \n",a)
}

func sUm(a int,b int) int {
	// fmt.Print(a+b)
	return a+b
}

func main() {
	// f1(111)

	sum:=sUm(10,20)
	fmt.Print(sum)
}