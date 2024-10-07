package main

import (
	"fmt"
)

const PI = 3.14 // Untyped constant
// const PI float32 = 3.14 // Typed constant

// Declare mutiple constants at once, no need to repeat keyword const and always untyped const
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

	DataTypes()
}

func DataTypes() {
	var a bool = true
	var b int = 5
	var c float32 = 3.14
	var d string = "Hi!"

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("String: ", d)
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
