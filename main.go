package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Flags
	showHelp := flag.Bool("h", false, "Show help information")
	decode := flag.Bool("d", false, "Decode")
	encode := flag.Bool("e", false, "Encode")
	multi := flag.Bool("m", false, "Multiple line toggle")
	server := flag.Bool("s", false, "Web server toggle")
	flag.Parse()

	//Server toggle
	if *server {
		runServer()
		return
	}

	//Show help if no arguments are found or -h flag is used
	if *showHelp || len(flag.Args()) == 0 {
		printHelp()
		return
	}

	//Decode by default
	if !*decode && !*encode {
		*decode = true
	}

	//Collect inputs
	var inputs [][]string
	for i := range flag.Args() {
		if strings.HasSuffix(flag.Args()[i], ".txt") {
			file, err := loadFile(flag.Args()[i])
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			inputs = append(inputs, file)
		} else {
			if strings.IndexAny(flag.Args()[i], "\n") != -1 {
				println("Invalid input detected in:", flag.Args()[i], " please use the -m toggle and follow the instructions")
				return
			}
			var line []string
			line = append(line, flag.Args()[i])
			inputs = append(inputs, line)
		}
	}

	var results [][]string
	var result []string
	var correct bool

	for i := range inputs {
		if len(inputs[i]) != 1 && *multi {
			if *encode {
				result, correct = encoder(inputs[i])
			} else {
				result, correct = decoder(inputs[i], flag.Args()[i])
			}
			results = append(results, result)
		} else if len(inputs[i]) > 1 {
			println("Multiple lines detected in file:", flag.Args()[i], " please turn on the -m toggle and follow the instructions")
		} else {
			if *encode {
				result, correct = encoder(inputs[i])
			} else {
				result, correct = decoder(inputs[i], flag.Args()[i])
			}
			results = append(results, result)
		}
		if !correct {
			if *server {

			} else {
				return
			}
		}
	}

	file, err := os.Create("latestOutput.txt")
	if err != nil {
		println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := range results {
		for j := range results[i] {
			_, err := writer.WriteString(results[i][j] + "\n")
			if err != nil {
				println("Error writing to file:", err)
				return
			}
		}

		_, err := writer.WriteString("\n")
		if err != nil {
			println("Error writing to file:", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		println("Error flushing to file:", err)
		return
	}

	//Print result
	for i := range results {
		for j := range results[i] {
			println(results[i][j])
		}
	}
}

func loadFile(path string) ([]string, error) {
	//Load file into a variable
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Create a slice of the lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
