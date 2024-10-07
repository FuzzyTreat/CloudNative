package main

import (
	"fmt"
)

func main() {
	// This is a comment
	fmt.Println("Hello World!")
	/* This is also a comment
	but on multiple rows */

	var explicitOne int32 = 0 // Full syntax declaration
	var typeLessOne = 0       // defaults to OS std byte size, ie int32 on a 32 bit system and int64 on a 64 bit system
	walrusType := 0           // <- Has to be assigned a value at declaration, can only be done in a function -> needs a scope
	var isRunning bool

	// declare multiple variables on a single row
	// NOTE: only same type of values if the datatype is specified
	// var a,b,c int32 = 1,3,7
	// var x,y,z = 1,"Fluff", 7

	fmt.Printf("Excplicit: %d", explicitOne)
	fmt.Printf("No type specified: %d", typeLessOne)
	fmt.Printf("Walrus: %d", walrusType)
	fmt.Printf("Default value: %v", isRunning)

	explicitOne = 42
	fmt.Printf("Altered value: %v", explicitOne)
}
