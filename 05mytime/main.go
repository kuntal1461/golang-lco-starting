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
	
}