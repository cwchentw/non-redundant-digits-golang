package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Create a scanner object
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt message
	fmt.Print("Please input a number: ")

	// Convert the input to number
	var num int
	if scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		num = n
	} else {
		log.Fatal("Failed reading from stdin")
	}

	count := 1
	for i := 1; i < int(math.Pow10(num)); i++ {
		// Convert the integer to string
		s := strconv.Itoa(i)

		// Create a map as a counter
		hash := make(map[rune]int)

		// Count digits
		for _, e := range s {
			hash[e] += 1
		}

		digit := 0
		j := i
		for j > 0 {
			j = j / 10
			digit += 1
		}

		if len(hash) >= digit {
			fmt.Printf("%d ", i)
			count += 1

			// Print EOL for each 10 numbers
			if count%10 == 0 {
				fmt.Println("")
			}
		}
	}

	// Print tailing EOL
	fmt.Println("")
}
