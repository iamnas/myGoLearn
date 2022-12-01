package main

import "fmt"

func main() {
	name,age := "nas",10 ;
	_,_=name,age; // multiple mute
	// fmt.Println(name,age) ;
	// // fmt.Print("name")
	// name,age1:="avi",11;
	// fmt.Println(name,age1) ;

	var (
		school string
		roll int 
		pass bool
	)

	fmt.Println(school,roll,pass)
	var x,y,z int
	fmt.Println(x,y,z)
}