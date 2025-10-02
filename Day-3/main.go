package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	message := "Welcome"
	fmt.Print(message, " \n")

	//getting input from the user 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating",input)

}
