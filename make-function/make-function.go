package main

import "fmt"

func main(){
	
	fmt.Println("Go Make Function Example")

	//Create a slice using the make function
	aSlice := make([]int,5)

	//Set and display a value in the slice
	aSlice[0] = 1

	fmt.Println(aSlice[0])

}
