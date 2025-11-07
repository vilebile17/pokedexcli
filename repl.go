package main

import "strings"

func CleanInput(text string) []string {
	newText := strings.ToLower(text)
	return strings.Fields(newText)	
}
