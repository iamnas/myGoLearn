package main

import "fmt"

func main() {
	name:= "aju"
	roll:=12
	_=name 
	// fmt.Println("My roll number is ",roll)

	// format specifier 
	// %d - integer
	// %s -String -
	// %f float

	fmt.Printf("My roll number is %d and name is %s\n",roll,name)
}