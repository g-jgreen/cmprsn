package main

import (
	"fmt"
)

func elementsInArray(checkElements []int, elements []int) int {
	var i int = 0
	var offset int = 0

	for element := range elements {

		if len(checkElements) <= offset {
			return i - len(checkElements)
		} else if checkElements[offset] == elements[element] {
			offset = offset + 1
		} else {
			offset = 0
		}
		i = i + 1
	}
	return -1
}

func main() {
	// const text string = "I AM SAM. I AM SAM. SAM I AM"
	const text string = "SAM SAM"
	textToBytes := []byte(text)

	var searchBuffer []int
	var checkChars []int
	var i int = 0

	for char := range textToBytes {
		checkChars = append(checkChars, int(textToBytes[char]))
		searchIndex := elementsInArray(checkChars, searchBuffer)

		if searchIndex == -1 || i == len(textToBytes)-1 {
			if len(checkChars) > 1 {
				index := elementsInArray(checkChars[:len(checkChars)-1], searchBuffer)
				offset := i - index - len(checkChars) + 1
				length := len(checkChars)

				t := fmt.Sprintf("<%v,%v>\n", offset, length)

				if len(t) > length {
					for _, num := range checkChars {
						fmt.Println(string(num))
					}

				} else {
					fmt.Println(t)
				}
			} else {
				fmt.Println(string(textToBytes[char]))
			}

			checkChars = nil
		}
		searchBuffer = append(searchBuffer, int(textToBytes[char]))

		i = i + 1
	}
}
