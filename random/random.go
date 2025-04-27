package random

import (
	"fmt"
	"math/rand"
)

func Random() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
