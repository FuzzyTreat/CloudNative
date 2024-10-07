package variables

import (
	"fmt"
)

func variables() {
	var explicitOne int32 = 0 // Full syntax declaration
	var typeLessOne = 0       // defaults to OS std byte size, ie int32 on a 32 bit system and int64 on a 64 bit system
	walrusType := 0           // <- Has to be assigned a value at declaration, can only be done in a function -> needs a scope
	var isRunning bool

	fmt.Printf("Excplicit: %d", explicitOne)
	fmt.Printf("No type specified: %d", typeLessOne)
	fmt.Printf("Walrus: %d", walrusType)
	fmt.Printf("Default value: %v", isRunning)

	explicitOne = 42
	fmt.Printf("Altered value: %v", explicitOne)
}
