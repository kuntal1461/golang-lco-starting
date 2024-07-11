package main

import (
	"fmt"
	"time"
)


func main()  {
	fmt.Println("Welcome to time study of golang")

	presentTime:= time.Now();
	fmt.Println(presentTime)

	// to get particular format 

	fmt.Println(presentTime.Format("01-02-2006"))

	fmt.Println("This is the year :",presentTime.Year())
	fmt.Println("This is the month:",presentTime.Month())
	fmt.Println("This is the day :", presentTime.Day())

	
}