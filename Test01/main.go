package main

import (
	"fmt"
)

const PI = 3.14 // Untyped constant
// const PI float32 = 3.14 // Typed constant

func main() {
	const LOCAL_PI = 3.15

	// This is a comment
	fmt.Println("Hello World!")
	/* This is also a comment
	but on multiple rows */

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
	fmt.Printf("Local Const value: %f\n", LOCAL_PI)

}
