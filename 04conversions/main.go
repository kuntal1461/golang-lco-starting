package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main()  {
	fmt.Println("Please rate our pizza")
	fmt.Println("The pizza rating will be in 1 to 5 ")

	redear:= bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the pizzar rating:  ")

	input,_:=redear.ReadString('\n')

	fmt.Println("Thanks of rating ,",input)


	numRating, err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 in your rating ",numRating+1)
	}


}