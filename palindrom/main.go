package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	runetext := []rune(text)
	len := len(runetext)
	for i := 0; ; i++ {
		if runetext[i] != runetext[len-(i+1)] {
			fmt.Println("Нет")
			break
		}
		if i == len/2 {
			fmt.Println("Палиндром")
			break
		}
	}
}
