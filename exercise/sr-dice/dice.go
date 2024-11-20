//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func rollDice(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	// numRolls := 10
	// numDice := 2
	// numSides := 6

	// for totalRoll := 0; totalRoll < numRolls; totalRoll++ {
	// 	rollSum := 0
	// 	for totalDice := 0; totalDice < numDice; totalDice++ {
	// 		rollSum += rand.Intn(numSides) + 1
	// 	}

	// 	fmt.Printf("Roll #%d: %d\n", totalRoll+1, rollSum)

	// 	if rollSum == 2 && numDice == 2 {
	// 		fmt.Println("Snake eyes")
	// 	} else if rollSum == 7 {
	// 		fmt.Println("Lucky 7")
	// 	} else if rollSum%2 == 0 {
	// 		fmt.Println("Even")
	// 	} else {
	// 		fmt.Println("Odd")
	// 	}
	// }

	dice, sides := 2, 6
	rolls := 3

	for totalRolls := 0; totalRolls < rolls; totalRolls++ {
		rollSum := 0
		for totalDice := 0; totalDice < dice; totalDice++ {
			rolled := rollDice(sides)
			rollSum += rolled
			fmt.Printf("Roll #%d: %d\n", totalRolls+1, rolled)
		}
		fmt.Println("Total rolled:", rollSum)

		switch sum := rollSum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake eyes")
		case sum == 7:
			fmt.Println("Lucky 7")
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}
}
