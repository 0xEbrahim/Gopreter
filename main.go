package main

import "fmt"

func main() {
	s := "😂"
	r := []rune(s)
	fmt.Printf("%c", r[0])
}
