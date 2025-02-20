package main

import (
	"strconv"
	"strings"
)

func encodeLine(input string) string {
	//Get the length of the input
	n := len(input)
	//If the input is empty, return an empty string
	if n == 0 {
		return ""
	}

	var builder strings.Builder

	/* -- Recursive Solution
	//Check for repeating substrings
	//l is the length of the substring
	for l := 1; l <= n/2; l++ {
		//i is the starting index of the substring
		for i := 0; i <= n-2*l; i++ {
			substr := input[i : i+l]
			count := 1
			//j is the starting index of the next substring
			for j := i + l; j <= n-l; j += l {
				//If the next substring is the same, increment the count
				if input[j:j+l] == substr {
					count++
				} else {
					//If the next substring is different, break the loop
					break
				}
			}
			//If the count is greater than 1 and is worth encoding, encode the substring
			if count+len(substr) > 6 && count > 1 {
				encoded := "[" + strconv.Itoa(count) + " " + substr + "]"
				builder.WriteString(input[:i])
				builder.WriteString(encoded)
				builder.WriteString(encodeLine(input[i+count*l:]))
				return builder.String()
			}
		}
	}
	//If no encoding is possible, return the input
	return input
	*/

	//Iterative Solution
	for i := 0; i < n; {
		encoded := false
		//Check for repeating substrings
		//l is the length of the substring
		for l := 1; l <= (n-i)/2; l++ {
			substr := input[i : i+l]
			count := 1
			//j is the starting index of the next substring
			for j := i + l; j <= n-l; j += l {
				//If the next substring is the same, increment the count
				if input[j:j+l] == substr {
					count++
				} else {
					//If the next substring is different, break the loop
					break
				}
			}
			//If the count is greater than 1 and is worth encoding, encode the substring
			if count+len(substr) > 6 && count > 1 {
				encoded = true
				encodedStr := "[" + strconv.Itoa(count) + " " + substr + "]"
				builder.WriteString(encodedStr)
				i += count * l
				break
			}
		}
		//If no encoding is possible, add the character to the result
		if !encoded {
			builder.WriteByte(input[i])
			i++
		}
	}
	return builder.String()
}

func encoder(input []string) ([]string, bool) {
	var result []string
	correct := true
	if input == nil {
		return nil, false
	}
	for _, line := range input {
		//Encode each line in the input
		encodedLine := encodeLine(line)
		result = append(result, encodedLine)
	}
	return result, correct
}
