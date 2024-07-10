package main

import (
	"bufio"
	"fmt"
	"os"
)



func main()  {
	welcome:= "Welcome to user input"
	fmt.Println(welcome)


	reader:= bufio.NewReader(os.Stdin);

	fmt.Println("please enter your age :")

	// comma ok syntex || err ok 

	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for etering your age ,",input)
	// want to know the data type of the input 

	fmt.Printf("The data type of the input is  %T",input)




}