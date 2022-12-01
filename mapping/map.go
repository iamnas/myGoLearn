package main

import "fmt"

func main() {
	// fmt.Print("hello")
	people:=map[string]int{}
	people["avi"] = 6

	students:=map[int]string{
		1:"avi",
		2:"nas",
	}
	// fmt.Print(students)

	for key,value := range students{
		fmt.Print(key," ",value," ")
	}

}