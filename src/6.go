
package main

import "fmt"

func getRandomNumber() int {
	return rand.Intn(10) + 1
}

func main() {
	randNumber := getRandomNumber()
	fmt.Println("Your random number is:", randNumber)
}