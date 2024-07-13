package main

import "fmt"



func main()  {
	fmt.Println("Welcome to pointer")

	// var myPtr *int
	// fmt.Println("The value of ptr: ",myPtr)

	myNumber :=23

	var myPtr = &myNumber

	// this is giving the memory value of this pointer 
	fmt.Println("the value of the pointer :",myPtr)
	// this is giving the actual value of this pointer 
	fmt.Println("the value of the pointer :", *myPtr)


	*myPtr = *myPtr+2

	fmt.Println("the sumation value is ",*myPtr)

}