package main

import "fmt"

const ( //Terminal color constants
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func printHelp() {
	//Help is printed when there are no arguments or -h flag is used
	fmt.Printf("\n%sArt Decoder / Encoder usage:%s\n", Blue, Reset)
	fmt.Printf("  go run . -h %s-- Show this help message%s\n", Red, Reset)
	fmt.Printf("  go run . %s-d/-e [-m]%s ./input.txt/\"string to work with\" %s-- Proper use of program%s\n", Yellow, Reset, Green, Reset)
	fmt.Printf("  go run . %s-s%s %s-- Proper use of program%s\n", Yellow, Reset, Green, Reset)
	fmt.Printf("\n%sDescription:%s\n", Yellow, Reset)
	fmt.Printf("  -d %s- Decodes your input%s\n", Yellow, Reset)
	fmt.Printf("  -e %s- Encodes your input%s\n", Yellow, Reset)
	fmt.Printf("  -m %s- Multiline mode%s\n", Yellow, Reset)
	fmt.Printf("  -s %s- Web server mode%s\n", Yellow, Reset)
	fmt.Println("  This program decodes or encodes your input.")
	fmt.Println("  This program decodes by default if there is no flag provided.")
	fmt.Println("  Arguments in square brackets are optional, this program is not encoding nor decoding input including square brackets")
	fmt.Println("  Providing invalid input will result in an error message")
	println("")
}
