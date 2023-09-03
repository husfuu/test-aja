package main

import (
	"fmt"
	"strings"
)

func GetManipulateName(name string) string {
	// input := "dani Maulana"
	// fmt.Println("input =", input)

	// Convert to lowercase and remove spaces
	name = strings.ReplaceAll(strings.ToLower(name), " ", "")

	mapInput := make(map[rune]int)
	var outputArr []string

	for _, v := range name {
		if _, ok := mapInput[v]; !ok {
			countChar := strings.Count(name, string(v))
			mapInput[v] = countChar
			if countChar == 1 {
				outputArr = append(outputArr, string(v))
			} else {
				outputArr = append(outputArr, fmt.Sprintf("%d%s", countChar, string(v)))
			}
		}
	}

	output := strings.Join(outputArr, "")
	return output
}
