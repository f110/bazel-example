package hello

import "fmt"

func Println(i, j int) {
	fmt.Printf("Hello %d\n", i+j)
}

func DebugPrint(i, j int) {
	Println(i, j)
}
