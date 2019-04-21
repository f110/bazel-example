package hello

import "fmt"

func Println(i int) {
	fmt.Printf("Hello %d\n", i)
}

func DebugPrint(i int) {
	Println(i)
}
