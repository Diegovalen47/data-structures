package basicmath

import "strconv"

func IsPalindrome(x int) bool {

	integerAsBytesArray := []byte(strconv.Itoa(x))
	var revertedBytesArray []byte

	for letter := len(integerAsBytesArray) - 1; letter >= 0; letter-- {
		revertedBytesArray = append(revertedBytesArray, integerAsBytesArray[letter])
	}

	return string(integerAsBytesArray) == string(revertedBytesArray)
}
