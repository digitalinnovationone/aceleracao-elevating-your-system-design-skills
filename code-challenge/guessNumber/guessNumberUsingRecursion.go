// Write a function that tries to guess the number 5 using math/rand lib.
// You can use from 1 to 10
// Make sure you have a limit of attempts
// Make it in a way that uses recursion

/*

In the provided code, the guessNumber function is called recursively.
 However, since there is a fixed limit on the number of attempts (attempts = 5000),
  the recursion is bounded and will not go infinitely deep. Therefore, the time complexity
  of the overall program is determined by the number of attempts, which is a constant value.

In this case, the time complexity is O(1) or constant time because the number of attempts is fixed at 5000,
 regardless of the input size. The recursion doesn't add to the time complexity as it has
 a fixed depth due to the limited number of attempts.

It's worth noting that if the number of attempts were variable and dependent on some input,
 the time complexity could be different. However, in this specific implementation, the number of attempts is constant.

 If we remove the limit on the number of attempts, the time complexity would be O(∞),
*/

package main

import (
	"fmt"
	"math/rand"
)

var numberToGuess = 15
var attempts = 5000
var attemptsCount = 0
var guessed = false

func guessNumber() bool {
	// Generate a random number
	randomNumber := rand.Intn(10000)
	attemptsCount++
	fmt.Println("Random number is: ", randomNumber, "Attempts: ", attemptsCount)

	if randomNumber == numberToGuess {
		fmt.Println("You guessed it!")
		return true
	}

	if attemptsCount < attempts {
		return guessNumber()
	}

	return false
}

func main() {
	fmt.Println("Hello, World!")

	guessNumber()

	if !guessed {
		fmt.Println("You did not guess it!")
	}
}
