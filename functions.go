package util

import (
	"fmt"

	common "github.com/ryhoh/go-util/common"
)

// Clips value into [min, max]
func Clip[T common.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

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

// Returns maximum element's index from parameters
func Argmax[T common.Ordered](param ...T) []int {
	switch len(param) {
	case 0:
		panic("Argmax(param) expects one or more parameters but given nothing")
	case 1:
		return []int{0}
	default:
		biggest := param[0]
		biggest_indices := []int{0}
		for i, rest := range param[1:] {
			if biggest < rest {
				biggest = rest
				biggest_indices = []int{i + 1}
			} else if biggest == rest {
				biggest_indices = append(biggest_indices, i+1)
			}
		}
		return biggest_indices
	}
}

// Returns maximum element's index from parameters
func Argmin[T common.Ordered](param ...T) []int {
	switch len(param) {
	case 0:
		panic("Argmin(param) expects one or more parameters but given nothing")
	case 1:
		return []int{0}
	default:
		smallest := param[0]
		smallest_indices := []int{0}
		for i, rest := range param[1:] {
			if smallest > rest {
				smallest = rest
				smallest_indices = []int{i + 1}
			} else if smallest == rest {
				smallest_indices = append(smallest_indices, i+1)
			}
		}
		return smallest_indices
	}
}

// Power for integer
// note: throws error when overflowed
func Pow[T1 common.Integer, T2 common.Unsigned](base T1, exponent T2) (T1, error) {
	if exponent == 0 {
		return 1, nil
	}
	if exponent == 1 {
		return base, nil
	}

	// 2 <= exponent
	res_A, err := Pow(base, exponent/2)
	if err != nil { // overflowed
		return base, err
	}

	res_B := res_A * res_A
	if (res_B / res_A) != res_A { // overflowed
		return base, fmt.Errorf("overflow occured")
	}

	if exponent%2 == 0 {
		return res_B, nil
	}

	res_C := res_B * base
	if (res_C / base) != res_B { // overflowed
		return base, fmt.Errorf("overflow occured")
	}
	return res_C, nil
}

// Reverses slice (in-place)
func Reverse[T interface{}](slice *[]T) {
	end := len(*slice)
	for i := 0; i < end/2; i++ {
		(*slice)[i], (*slice)[end-1-i] = (*slice)[end-1-i], (*slice)[i]
	}
}

// Returns reversed slice (returns copy and doesn't modify original)
func Reversed[T interface{}](slice *[]T) *[]*T {
	end := len(*slice)
	res := make([]*T, end)
	for i := 0; i < end; i++ {
		res[i] = &(*slice)[end-1-i]
	}
	return &res
}

// Reverse string
func ReversedString(str string) string {
	runes := []rune(str)
	Reverse(&runes)
	return string(runes)
}
