package main

import (
	"fmt"
	"math/rand"
)

var MAX = 100
var LOW = 1
var FIRST = true
var USR_ANSWER string
var COUNT = 0
var GUESS int

func get_input(){
	if(COUNT == 0){
		COUNT+=1
		fmt.Print("Pick a number between 1-100 and I will make a guess.\nEnter h for too high l for too low y if correct: ")
	} else{
		fmt.Print("Enter(h/l/y):")
	}
	fmt.Scan(&USR_ANSWER)
}


func run() {
	for USR_ANSWER != "y"{
		if FIRST{
			FIRST=false
			GUESS = MAX/2
		}else {
			if MAX==LOW{
				GUESS = MAX
				break
			}
			GUESS = rand.Intn(MAX - LOW) + LOW
		}
		fmt.Printf("Is your number %d?", GUESS)
		get_input()
		switch USR_ANSWER {
			case "h":
				MAX = GUESS - 1
			case "l":
				LOW = GUESS+1
			case "y":
				break
		}
	}
	fmt.Printf("Your Number %d Was Easy!!!\nMake it Harder Next Time.", GUESS)
}

// main function
func main() {
	run()
}