package util

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	/* Parameter changing */
	if actual := Max(1); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	if actual := Max(1, 2); actual != 2 {
		t.Errorf("expected 2 but given %d", actual)
	}

	if actual := Max(1, 2, 3); actual != 3 {
		t.Errorf("expected 3 but given %d", actual)
	}

	if actual := Max(3, 2, 1); actual != 3 {
		t.Errorf("expected 3 but given %d", actual)
	}

	/* Slice deploying */
	s := []int64{2, 3, 1}
	if actual := Max(s...); actual != 3 {
		t.Errorf("expected 3 but given %d", actual)
	}

	/* Error case */
	defer func() {
		err := recover()
		if err != "Max(param) expects one or more parameters but given nothing" {
			t.Errorf("unhandled error thrown: %v", err)
		}
	}()
	Max[int]()
}

func TestMin(t *testing.T) {
	/* Parameter changing */
	if actual := Min(1); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	if actual := Min(1, 2); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	if actual := Min(1, 2, 3); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	if actual := Min(3, 2, 1); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	/* Slice deploying */
	s := []int64{2, 3, 1}
	if actual := Min(s...); actual != 1 {
		t.Errorf("expected 1 but given %d", actual)
	}

	/* Error case */
	defer func() {
		err := recover()
		if err != "Min(param) expects one or more parameters but given nothing" {
			t.Errorf("unhandled error thrown: %v", err)
		}
	}()
	Min[int]()
}

func TestReverse(t *testing.T) {
	/* Length: odd or even */
	expected := []int{5, 4, 3, 2, 1}
	actual := []int{1, 2, 3, 4, 5}
	Reverse(&actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = []int{9, 8, 7, 6}
	actual = []int{6, 7, 8, 9}
	Reverse(&actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	/* non-ascii charactors */
	expected_rune := []rune("参弐壱")
	actual_rune := []rune("壱弐参")
	Reverse(&actual_rune)
	if !reflect.DeepEqual(actual_rune, expected_rune) {
		t.Errorf("expected %v but given %v", actual_rune, expected_rune)
	}
}

func TestReversedString(t *testing.T) {
	if actual := ReversedString("qwerty"); actual != "ytrewq" {
		t.Errorf("expected \"ytrewq\" but given %s", actual)
	}
	if actual := ReversedString("＼(^o^)／"); actual != "／)^o^(＼" {
		t.Errorf("expected \"／(^o^)＼\" but given %s", actual)
	}
}

/* Benchmarks */
func BenchmarkReverse(b *testing.B) {
	base := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		base[i] = i
	}
	b.ResetTimer()
	Reverse(&base)
}
