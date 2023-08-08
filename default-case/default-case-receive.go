package main

import (
	"fmt"
	"time"
)

func main(){
	
	fmt.Println("Default case when receiving Example")

	//Create two channels to use later with the select statement
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Go routine to send to first channel
	go func(){
		channel1 <- 65535
	}()

	//Go routine to send to second channel
	go func() {
		channel2 <- 12345
	}()

	//Select statement that receives values from the two channels, default should usually run
	select{
		case value := <-channel1:
			fmt.Println(value)
		case value := <-channel2:
			fmt.Println(value)
		default:
			fmt.Println("None ready!")
			time.Sleep(time.Second * 1)
	}

	//Give time for one of the goroutines to finish
	time.Sleep(time.Second * 1)

}
