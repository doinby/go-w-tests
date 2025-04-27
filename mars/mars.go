package mars

import "fmt"

func Mars() {
	// fmt.Print("My weight on the surface of Mars is ")
	// fmt.Print(149 * 0.3783)
	// fmt.Println(" lbs, and I would be ")
	// fmt.Print(32 * 365 / 687)
	// fmt.Print(" years old.")
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 32*365/687)
}
