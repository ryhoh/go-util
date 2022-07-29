package util

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func SetFromSlice[T comparable](slice *[]T) *Set[T] {
	res := Set[T]{}
	for _, value := range *slice {
		res.Add(value)
	}
	return &res
}

func (set *Set[T]) Add(value T) {
	(*set)[value] = struct{}{}
}

func (set *Set[T]) AddAll(values *[]T) {
	for _, value := range *values {
		(*set)[value] = struct{}{}
	}
}

func (set *Set[T]) Remove(value T) {
	delete(*set, value)
}

func (set *Set[T]) RemoveAll(values *[]T) {
	for _, value := range *values {
		delete(*set, value)
	}
}

func (set *Set[T]) Contains(value T) bool {
	for key := range *set {
		if value == key {
			return true
		}
	}
	return false
}

func (set *Set[T]) String() string {
	res := strings.Builder{}
	i := 0
	length := len(*set)

	res.WriteString("Set[")
	for key := range *set {
		res.WriteString(fmt.Sprint(key))
		if i != length-1 {
			res.WriteString(", ")
		}
		i++
	}

	res.WriteString("]")
	return res.String()
}

func Push[T any](slice *[]T, value ...T) {
	slice_new := append(*slice, value...)
	*slice = slice_new
}

func Pop[T any](slice *[]T) *T {
	if len(*slice) == 0 {
		return nil
	}

	end_idx := len(*slice) - 1
	retval := (*slice)[end_idx]
	slice_new := (*slice)[:end_idx]
	*slice = slice_new
	return &retval
}

func Top[T any](slice *[]T) *T {
	if len(*slice) == 0 {
		return nil
	}
	end_idx := len(*slice) - 1
	retval := (*slice)[end_idx]
	return &retval
}

func Enqueue[T any](slice *[]T, value ...T) {
	slice_new := append(*slice, value...)
	*slice = slice_new
}

func Dequeue[T any](slice *[]T) *T {
	if len(*slice) == 0 {
		return nil
	}

	retval := (*slice)[0]
	slice_new := (*slice)[1:]
	*slice = slice_new
	return &retval
}

func Front[T any](slice *[]T) *T {
	if len(*slice) == 0 {
		return nil
	}
	retval := (*slice)[0]
	return &retval
}
