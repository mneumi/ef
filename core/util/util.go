package util

import "strings"

func SnakeNameToPascal(snakeName string) string {
	wordSlice := strings.Split(snakeName, "_")

	for i := range wordSlice {
		wordSlice[i] = upperFirstLetter(wordSlice[i])
	}

	return strings.Join(wordSlice, "")
}

func SnakeNameToCamelName(snakeName string) string {
	wordSlice := strings.Split(snakeName, "_")

	for i := range wordSlice {
		if i != 0 {
			wordSlice[i] = upperFirstLetter(wordSlice[i])
		}
	}

	return strings.Join(wordSlice, "")
}

func upperFirstLetter(str string) string {
	byteSlice := []byte(str)
	if byteSlice[0] >= 97 && byteSlice[0] <= 122 {
		byteSlice[0] = byteSlice[0] - 32
	}
	return string(byteSlice)
}
