package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		value := int(i)
		fmt.Println(&value)
		fmt.Println(&i)
		fmt.Println(i)
		fmt.Println(value)
	}
	fmt.Println("hypotesis succeded it does work as a bypass to same pointer")
}
