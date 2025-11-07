package main

import (
	"bufio"
	"os"
	"fmt"
) 

func main() {
	s := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if inputExists := s.Scan(); inputExists {
			input := s.Text()
			bong := CleanInput(input)
			fmt.Printf("Your command was: %v\n", bong[0])
		}
	}
}
