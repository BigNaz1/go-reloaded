package main

import (
	"fmt"
	"os"
	"reloaded"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	modifiedText := reloaded.Atoan(string(text))
	modifiedText = reloaded.ReplaceSymbol(modifiedText)
	modifiedText = reloaded.HexatoDec(modifiedText)
	modifiedText = reloaded.BintoDec(modifiedText)
	modifiedText = reloaded.Capitalize(modifiedText)
	modifiedText = reloaded.ToUpper(modifiedText)
	modifiedText = reloaded.ToLower(modifiedText)
	modifiedText = reloaded.TransformText(modifiedText)
	modifiedText = reloaded.AdjustPunctuationSpacings(modifiedText)
	modifiedText = reloaded.RemoveExtraPunc(modifiedText)
	modifiedText = reloaded.AdvancedPunctuation(modifiedText)

	
	err = os.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Println("Text processing complete. Check your result file or view in terminal: " + string(modifiedText))
}
