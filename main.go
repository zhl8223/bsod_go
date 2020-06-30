package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var dll NtDLL
	dll.init()
	dll.bsod()
}
