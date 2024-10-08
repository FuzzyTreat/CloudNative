package main

import "fmt"

func main() {
	// GetLargestNumber()
	// GetLongestName()

	BankAccount()

}

func BankAccount() {
	// Map

	// Skriv en bankapplikation. Skriv först en meny med följande val:

	// Skapa konto
	// Ta bort konto
	// Lista alla kontonummer
	// Lista totalsaldo
	// Lista alla kontonummer och saldo

	// Felhantering behöver du inte bry dig om i steg 1.
}

func GetLongestName() {
	var userStrings = []string{}
	var inputString string
	var longestString string = ""

	for {
		fmt.Print("Enter a name: ")
		fmt.Scan(&inputString)
		fmt.Println()

		// Add the string to slice
		userStrings = append(userStrings, inputString)

		// compare input with current longest string
		if len(inputString) > len(longestString) {
			longestString = inputString
		}

		fmt.Print("Enter another name? (y/n): ")
		fmt.Scan(&inputString)
		fmt.Println()

		if inputString == "n" {
			break
		}
	}

	fmt.Printf("Longest string was: %s\n", longestString)

	fmt.Printf("The entered strings was\n")
	for _, name := range userStrings {
		fmt.Printf("Name: %s\n", name)
	}
}

func GetLargestNumber() {
	var selected int
	var idx int = 0
	var largest int = 0

	var userNumbers = [4]int{0, 0, 0, 0}

	for idx = range userNumbers {
		fmt.Print("Enter a number: ")
		fmt.Scan(&selected)
		userNumbers[idx] = selected
		fmt.Println()
	}

	for _, number := range userNumbers {
		if number > largest {
			largest = number
		}
	}

	fmt.Printf("Largest number entered was: %d\n", largest)
}
