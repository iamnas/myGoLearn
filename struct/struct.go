package main

import "fmt"

func main() {
	type Book struct{
		name string 
		id uint
		author string
	}

	myBook :=Book{"AAV",11,"ZMAQ"}
	fmt.Print(myBook)
}