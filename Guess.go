package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Random number genarator with UnixTime
	genarator := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(genarator)
	number := randomizer.Intn(10)

	var guess int
	num := 1
	for {

		fmt.Printf("\nWelcome, Guess a number from 1 to 10\n")
		fmt.Printf("This is you number %d attempt : \n", num)
		fmt.Scan(&guess)

		if guess < number {
			fmt.Println("Your guess is Low, Try a higher value")
			num++
		} else if guess > number {
			fmt.Println("Your guess is High, Try a smaller value")
			num++
		} else {
			fmt.Printf("Bingo! Correct guess on your %d attempt\n", num)
			break
		}

	}
}
