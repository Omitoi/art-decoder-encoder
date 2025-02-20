package main

import (
	"regexp"
	"strconv"
	"strings"
)

func decoder(input []string, fileName string) ([]string, bool) {
	fail := false
	//Define the regular expression for encoded substrings
	re := regexp.MustCompile(`\[(\d+) (.+?)\]`)
	//Find and replace all encoded substrings
	for i, str := range input {
		input[i] = re.ReplaceAllStringFunc(str, func(match string) string {
			submatches := re.FindStringSubmatch(match)
			//Convert the number of times to repeat the characters to an integer
			times, err := strconv.Atoi(submatches[1])
			//If the conversion fails, print an error and exit
			if err != nil {
				println("One of brackets started with a non-number character on line:", i+1, "in "+fileName)
				fail = true
			}
			//Repeat the characters the specified number of times
			characters := submatches[2]
			return strings.Repeat(characters, times)
		})
		if fail {
			return nil, false
		}
	}

	//Check for errors in the input

	//Check for brackets that are not separated by a space
	iCheck := regexp.MustCompile(`\[(\d+)(.+?)\]`)
	//Check for brackets that have a second argument empty
	iCheck2 := regexp.MustCompile(`\[(\d+)( )\]`)
	//Check for brackets that start with a non-number character
	iCheck3 := regexp.MustCompile(`\[\D`)

	for i, str := range input {
		if iCheck3.MatchString(str) {
			println("One of brackets started with a non-number character on line:", i+1, "in "+fileName)
			return nil, false
		}
		if iCheck2.MatchString(str) {
			println("One of brackets had a second argument empty on line:", i+1, "in "+fileName)
			return nil, false
		}
		if iCheck.MatchString(str) {
			println("One of brackets wasn't separated by a space on line:", i+1, "in "+fileName)
			return nil, false
		}
		//Check for unmatched brackets
		if strings.Contains(str, "[") || strings.Contains(str, "]") {
			println("Unmatched brackets detected on line:", i+1, "in "+fileName)
			return nil, false
		}
	}

	return input, true
}
