package main

import "fmt"

func main(){
	
	fmt.Println("For loop channel Example (No Range)")

	//Create a buffered channel with a capacity of 3
	bufChannel := make(chan int, 3)

	//Send three values to the channel
	for i := 0; i < 3; i++ {
		bufChannel <- i
	}

	//Uncomment this line for closer alternative to range
	//close(bufChannel)

	//Receive three values from the channel
	/*var x int //uncomment these lines for closer alternative to range
	for ok := true; ok;{*/
	for i := 0; i < 3; i++ {
	        /*x, ok = <-bufChannel //uncomment these lines
	        if !ok {
	                break
	        }
	        fmt.Println(x)*/
		fmt.Println(<-bufChannel)
	}

}
