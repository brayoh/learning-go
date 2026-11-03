// Formatting string in GO using various formatting verbs

package main

import "fmt"

type Area struct {
	width, height int
}

func main() {
	area := Area{width: 30, height: 40}
	isActive := true

	// --- 1. Struct & Meta Verbs ---
	// Struct-aware %+v (Includes keys)
	fmt.Printf("1. Struct values (%%+v): %+v \n", area)

	// Go-syntax %#v (Re-constructible Go code representation)
	fmt.Printf("2. GO syntax (%%#v): %#v\n", area)

	// Type checking
	fmt.Printf("3. Struct Type (%%T): %T \n\n", area)

	// --- 2. String & Boolean Verbs ---
	var areaFloat = area.calculateArea("This is a receiver function parameter")
	fmt.Printf("4. Unquoted string (%%s): %s \n", "Hello Go!")
	fmt.Printf("5. Quoted string (%%q): %q \n", "Hello Go!")
	fmt.Printf("6. Boolean value (%%t): %t \n\n", isActive)

	// --- 3. Floating-Point Verbs ---
	fmt.Printf("7. Raw float value (%%v): %v \n", areaFloat)
	fmt.Printf("8. Float with standard decimal (%%f): %f \n", areaFloat)
	fmt.Printf("9. Float with 2 decimal places (%%.2f): %.2f \n", areaFloat)
	fmt.Printf("10. Scientific notation (%%e): %e \n\n", areaFloat)

	// --- 4. Integer Base Verbs (using the width property) ---
	fmt.Printf("11. Integer Base 10 (%%d): %d \n", area.width)
	fmt.Printf("12. Integer Binary (%%b): %b \n", area.width)
	fmt.Printf("13. Integer Hexadecimal (%%x): %x \n", area.width)
}

func (a Area) calculateArea(log string) float64 {
	fmt.Printf("--- Inside Receiver Function ---\n")
	fmt.Printf("Log message: %s \n", log)
	fmt.Printf("-------------------------------\n")

	return float64(a.height) * float64(a.width)
}
