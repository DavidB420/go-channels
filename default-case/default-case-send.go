package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Default case when sending Example")

	//Create two channels to use later with the select statement
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Go routine to receive from first channel
	go func(){
		value := <-channel1
		fmt.Println(value)
	}()

	//Go routine to receive from second channel
	go func(){
		value := <-channel2
		fmt.Println(value)
	}()

	//Select statement that sends value to the two channels, default should usually run
	select{
		case channel1 <- 65535:
		case channel2 <- 12345:
		default:
			fmt.Println("None ready!")
			time.Sleep(time.Second * 1)
	}
	
	//Give time for one of the goroutines to finish
	time.Sleep(time.Second * 1)

}
