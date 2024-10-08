package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PI = 3.14 // Untyped constant
// const PI float32 = 3.14 // Typed constant

// Declare mutipleconstants at once, no need to repeat keyword const and always untyped const
// No complaint about un-used constants
const (
	A int = 11
	B     = 3.16
	C     = "Hello there!"
)

func main() {
	const LOCAL_PI = 3.17 // tied to the current scope

	// This is a comment
	fmt.Println("Hello World!")
	/* This is also a comment
	but on multiple rows */

	// Variables()
	// fmt.Printf("Main Local Const value: %f\n", LOCAL_PI)

	//Output()
	//DataTypes()
	StringValues()
}

func StringValues() {
	// String #1
	userStrings := []string{}
	var inputString string

	fmt.Printf("Enter 3 strings. \n")

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter string %d: ", i+1)
		fmt.Scan(&inputString)
		userStrings = append(userStrings, inputString)
		fmt.Printf("\n")
	}

	output := strings.Join(userStrings, " ")
	// fmt.Printf("%s %s %s\n", userStrings[0], userStrings[1], userStrings[2])
	fmt.Println(output)

	// String #2
	firstIdx := strings.Index(C, "t")
	lastIdx := strings.LastIndex(C, "o")

	fmt.Printf("\n")
	fmt.Printf("First index of t: %d in %s\n", firstIdx, C)
	fmt.Printf("Last index of o: %d in %s\n", lastIdx, C)

	// string #3
	fullName := "kurt andersson"
	fmt.Printf("Name before formatting: %s\n", fullName)

	fullName = strings.Replace(fullName, string(fullName[0]), strings.ToUpper(string(fullName[0])), 1)
	spaceIdx := strings.Index(fullName, " ")
	fullName = strings.Replace(fullName, string(fullName[spaceIdx+1]), strings.ToUpper(string(fullName[spaceIdx+1])), 1)

	fmt.Printf("Formatted name: %s\n", fullName)

	// string #4
	exampleString := "Detta är en sträng som du skall ändra"
	exampleString = strings.ReplaceAll(exampleString, " ", "*")
	fmt.Printf("Number of stars after replacement: %d\n", strings.Count(exampleString, "*"))

	// string #5
	fmt.Print("\nStrings #5\n")
	var email string
	var isCorrect bool
	var suffixCount int16 = 0

	fmt.Printf("Enter an email address: ")
	fmt.Scan(&email)
	fmt.Println(" ")

	atIdx := strings.Index(email, "@")
	dotIdx := strings.LastIndex(email, ".")

	if atIdx < 0 || dotIdx < 0 {
		isCorrect = false
	} else {
		for i := dotIdx; dotIdx < len(email)-dotIdx; i++ {
			suffixCount = suffixCount + 1
		}

		if suffixCount < 2 {
			isCorrect = false
		} else {
			isCorrect = true
		}
	}

	if isCorrect {
		fmt.Println("You entered a valid email address!")
	} else {
		fmt.Println("You entered a invalid email address!")
	}

	Strings6()
}

func Strings6() {
	// string #6
	var text string
	var reverseText string
	fmt.Print("\nStrings #6\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()

		if len(text) > 0 {
			break
		}
	}

	for i := len(text); i > 0; i-- {
		reverseText += string(text[i-1])
	}

	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))
	reverseText = strings.ToLower(strings.ReplaceAll(reverseText, " ", ""))

	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Reverse: %s\n", reverseText)

	if text == reverseText {
		fmt.Printf("The text is a palindrom!\n")
	} else {
		fmt.Printf("The text is not a palindrom!\n")
	}
}

func DataTypes() {

	var stop bool = false
	var menuChoice int32
	var accountName string
	var amount int

	accountList := make(map[string]int)

	for {

		fmt.Println("")
		fmt.Println("1. Skapa konto")
		fmt.Println("2. Ta bort konto")
		fmt.Println("3. Lista alla konton")
		fmt.Println("4. Visa totalsaldo")
		fmt.Println("5. Lista all konton och saldonn")
		fmt.Println("6. Avsluta")
		fmt.Println("")

		fmt.Scan(&menuChoice)

		switch menuChoice {
		case 1:
			{
				fmt.Println("")
				fmt.Println("Enter account name: ")
				fmt.Scan(&accountName)
				fmt.Println("")
				fmt.Println("Enter amount: ")
				fmt.Scan(&amount)
				fmt.Println("")

				accountList[accountName] = amount
				stop = false
			}
		case 2:
			{
				fmt.Println("")
				fmt.Println("Enter account name: ")
				fmt.Scan(&accountName)
				fmt.Println("")

				delete(accountList, accountName)
				fmt.Println("")
				fmt.Printf("Account %s has been deleted.\n", accountName)
				stop = false
			}
		case 3:
			{
				for key, _ := range accountList {
					fmt.Printf("Account name: %s.\n", key)
				}

				stop = false
			}
		case 4:
			{
				var total int

				for _, value := range accountList {
					total = total + value
				}

				fmt.Println("")
				fmt.Printf("Totalt saldo: %d\n", total)

				stop = false
			}
		case 5:
			{
				{
					for key, value := range accountList {
						fmt.Printf("Account name: %s. Saldo: %d\n", key, value)
					}

					stop = false
				}
			}
		case 6:
			{
				stop = true
			}
		default:
		}

		if stop {
			break
		}
	}
}

func Output() {
	var hello = "Hello"
	var world string = "world!"
	var x, y = 10, 20
	var hexValue = 65056

	fmt.Println("Print with no formatting")
	fmt.Print(hello)
	fmt.Print(world)
	fmt.Println("")

	fmt.Print("Print with new line char\n")
	fmt.Print(hello, "\n")
	fmt.Print(world, "\n")

	fmt.Println("Output two values in one statement")
	fmt.Print(hello, "\n", world)
	fmt.Println("")

	fmt.Println("Output two int values in one statement, no formatting between them")
	fmt.Print(x, y)
	fmt.Println("")

	fmt.Println("Output formatted text")
	fmt.Printf("This text has beeen formatted with %s and %s\n", hello, world)

	fmt.Printf("Const PI default format: %v\n", PI)
	fmt.Printf("Const PI Go syntax format: %#v\n", PI)
	fmt.Printf("Const PI type format: %T\n", PI)
	fmt.Printf("Print the '%%', uses double %% as formatter\n")

	fmt.Printf("Double quted string, no extra quotes needed: %q\n", hello)
	fmt.Printf("Integer as a Hex dump byte values: %x\n", hexValue)

	fmt.Printf("Output float value: %f\n", PI)
	fmt.Printf("Output float value, 2 decimals: %.2f\n", PI)
	fmt.Printf("Output float exponent as needed: %g\n", PI)
}

func Variables() {
	const LOCAL_PI = 3.15 // tied to the current scope, not the same as in main

	var explicitOne int32 = 1 // Full syntax declaration
	var typeLessOne = 2       // defaults to OS std byte size, ie int32 on a 32 bit system and int64 on a 64 bit system
	walrusType := 3           // <- Has to be assigned a value at declaration, can only be done in a function -> needs a scope
	var isRunning bool

	// declare multiple variables on a single row
	// NOTE: only one type of values if the datatype is specified
	// var a,b,c int32 = 1,3,7
	// var x,y,z = 1,"Fluff", 7

	fmt.Printf("Explicit: %d\n", explicitOne)
	fmt.Printf("No type specified: %d\n", typeLessOne)
	fmt.Printf("Walrus: %d\n", walrusType)
	fmt.Printf("Default value: %v\n", isRunning)

	explicitOne = 42
	fmt.Printf("Altered value: %v\n", explicitOne)

	// Camel case: myVariableName = "One"
	// Pascal case: MyVariableName = "One" <- Typical for C# etc
	// Snake case: my_variable_name = "One"

	fmt.Printf("Global Const value: %f\n", PI)
	fmt.Printf("Func Const value: %f\n", LOCAL_PI)
}
