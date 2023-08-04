package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	
	fmt.Println("Unbuffered Go channel Example")

	//Make an unbuffered channel 
	unbufChannel := make (chan int)
	
	//Make a synchronous wait group
	var wg sync.WaitGroup
	wg.Add(2)

	//Sender goroutine
	go func(){
	
		defer wg.Done() //Mark as done

		unbufChannel <- 65535
	}()

	//Receiver goroutine
	go func(){
		defer wg.Done() //Mark as done

		//Add a delay to demonstrate the blocking
		time.Sleep(time.Second * 3)
		
		//Output the received message
		fmt.Println(<-unbufChannel)
		
	}()
	
	//Wait for both routines to be marked as done
	wg.Wait()
}
