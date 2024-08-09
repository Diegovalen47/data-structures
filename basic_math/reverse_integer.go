package basicmath

import (
	"errors"
	"strconv"
)

func Reverse(x int) int {

	signOperator := ternary(x < 0, -1, 1).(int)

	integerAsBytesArray := []byte(strconv.Itoa(x * signOperator))
	var revertedBytesArray []byte

	for letter := len(integerAsBytesArray) - 1; letter >= 0; letter-- {
		revertedBytesArray = append(revertedBytesArray, integerAsBytesArray[letter])
	}

	integerReverted, err := strToInt32(string(revertedBytesArray))

	if err != nil {
		return 0
	}

	return integerReverted * signOperator

}

func ternary(condition bool, valueIfTrue, valueIfFalse interface{}) interface{} {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}

func strToInt32(str string) (int, error) {

	integer, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, errors.New("entero mayor a 32 bits")
	}

	result := int(integer)

	return result, err
}
