package main

import (
	"runtime"
)

func main() {
	var N = 40000000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i) //is it a way to bypass double pointer?
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
}
