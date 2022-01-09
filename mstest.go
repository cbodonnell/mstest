package mstest

import (
	"strconv"
)

func Solution(S string, T string) bool {
	// write your code in Go 1.4
	sSlice := StringToLetterSlice(S)
	tSlice := StringToLetterSlice(T)
	// first check if they have the same number of letters
	if len(sSlice) != len(tSlice) {
		return false
	}
	// array are the same length, now check
	// if they could be the same word
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] != "" && tSlice[i] != "" && sSlice[i] != tSlice[i] {
			// letters are different, not same word
			return false
		}
	}
	// all letters are same or could possibly be the same
	return true
}

func StringToLetterSlice(s string) []string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		_, err := strconv.Atoi(string(s[i]))
		if err != nil {
			// current char is not a number
			// add it to the slice
			result = append(result, string(s[i]))
			continue
		}
		// current char at position i is a number
		numString := string(s[i])
		// need to check if there are more numbers
		for j := i + 1; j < len(s); j++ {
			_, err := strconv.Atoi(string(s[j]))
			if err != nil {
				// current char is not a number
				break
			}
			// next char is also a number
			// append to the current numString
			numString = numString + string(s[j])
			// also need to increment i
			// to avoid double counting
			i++
		}
		num, _ := strconv.Atoi(numString)
		for k := 0; k < num; k++ {
			result = append(result, "")
		}
	}
	return result
}
