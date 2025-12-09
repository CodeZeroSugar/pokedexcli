package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		firstWord := strings.Fields(strings.ToLower(strings.TrimSpace(scanner.Text())))[0]
		fmt.Printf("Your command was: %v\n", firstWord)
	}
}
