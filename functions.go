package util

import (
	common "github.com/ryhoh/go-util/common"
)

// Returns maximum element from parameters
func Max[T common.Ordered](param ...T) T {
	switch len(param) {
	case 0:
		panic("Max(param) expects one or more parameters but given nothing")
	case 1:
		return param[0]
	default:
		biggest := param[0]
		for _, rest := range param[1:] {
			if biggest < rest {
				biggest = rest
			}
		}
		return biggest
	}
}

// Returns maximum element from parameters
func Min[T common.Ordered](param ...T) T {
	switch len(param) {
	case 0:
		panic("Min(param) expects one or more parameters but given nothing")
	case 1:
		return param[0]
	default:
		smallest := param[0]
		for _, rest := range param[1:] {
			if smallest > rest {
				smallest = rest
			}
		}
		return smallest
	}
}

// Reverses slice (in-place)
func Reverse[T interface{}](slice *[]T) {
	end := len(*slice)
	for i := 0; i < end/2; i++ {
		(*slice)[i], (*slice)[end-1-i] = (*slice)[end-1-i], (*slice)[i]
	}
}

func ReversedString(str string) string {
	runes := []rune(str)
	Reverse(&runes)
	return string(runes)
}
