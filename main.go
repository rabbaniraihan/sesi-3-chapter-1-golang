package main

import "fmt"

func main() {
	var word string = "Selamat malam"
	for i := 0; i <= len(word)-1; i++ {
		fmt.Println(string(word[i]))
	}

	count := make(map[string]int)
	for _, char := range word {
		count[string(char)]++
	}

	fmt.Println(count)
}
