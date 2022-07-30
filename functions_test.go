package util

import (
	"reflect"
	"testing"
)

func TestClip(t *testing.T) {
	/* uint8 [0, 255] clipping */
	if actual := Clip(20, 0, 255); actual != 20 {
		t.Errorf("expected %v but given %v", 20, actual)
	}
	if actual := Clip(-20, 0, 255); actual != 0 {
		t.Errorf("expected %v but given %v", 0, actual)
	}
	if actual := Clip(400, 0, 255); actual != 255 {
		t.Errorf("expected %v but given %v", 255, actual)
	}
}

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

func TestArgmax(t *testing.T) {
	/* Parameter changing */
	if actual := Argmax(1); !reflect.DeepEqual(actual, []int{0}) {
		t.Errorf("expected [0] but given %d", actual)
	}

	if actual := Argmax(1, 2); !reflect.DeepEqual(actual, []int{1}) {
		t.Errorf("expected [1] but given %d", actual)
	}

	if actual := Argmax(1, 2, 3); !reflect.DeepEqual(actual, []int{2}) {
		t.Errorf("expected [2] but given %d", actual)
	}

	if actual := Argmax(3, 2, 3); !reflect.DeepEqual(actual, []int{0, 2}) {
		t.Errorf("expected [0, 2] but given %d", actual)
	}

	/* Slice deploying */
	s := []int64{2, 3, 1}
	if actual := Argmax(s...); !reflect.DeepEqual(actual, []int{1}) {
		t.Errorf("expected [1] but given %d", actual)
	}

	/* Error case */
	defer func() {
		err := recover()
		if err != "Argmax(param) expects one or more parameters but given nothing" {
			t.Errorf("unhandled error thrown: %v", err)
		}
	}()
	Argmax[int]()
}

func TestArgmin(t *testing.T) {
	/* Parameter changing */
	if actual := Argmin(1); !reflect.DeepEqual(actual, []int{0}) {
		t.Errorf("expected [0] but given %d", actual)
	}

	if actual := Argmin(2, 1); !reflect.DeepEqual(actual, []int{1}) {
		t.Errorf("expected [1] but given %d", actual)
	}

	if actual := Argmin(3, 2, 1); !reflect.DeepEqual(actual, []int{2}) {
		t.Errorf("expected [2] but given %d", actual)
	}

	if actual := Argmin(1, 2, 1); !reflect.DeepEqual(actual, []int{0, 2}) {
		t.Errorf("expected [0, 2] but given %d", actual)
	}

	/* Slice deploying */
	s := []int64{2, 1, 3}
	if actual := Argmin(s...); !reflect.DeepEqual(actual, []int{1}) {
		t.Errorf("expected [1] but given %d", actual)
	}

	/* Error case */
	defer func() {
		err := recover()
		if err != "Argmin(param) expects one or more parameters but given nothing" {
			t.Errorf("unhandled error thrown: %v", err)
		}
	}()
	Argmin[int]()
}

func TestAll(t *testing.T) {
	/* Parameter changing */
	expected := false
	if actual := All(); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = true
	if actual := All(true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = false
	if actual := All(false, true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = false
	if actual := All(false, false, false); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = false
	if actual := All(false, false, true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = true
	if actual := All(true, true, true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	/* Slice deploying */
	s := []bool{false, false, true}
	expected = false
	if actual := All(s...); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}
}

func TestAny(t *testing.T) {
	/* Parameter changing */
	expected := false
	if actual := Any(); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = true
	if actual := Any(true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = true
	if actual := Any(false, true); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	expected = false
	if actual := Any(false, false, false); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	/* Slice deploying */
	s := []bool{false, false, true}
	expected = true
	if actual := Any(s...); actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}
}

func TestPow(t *testing.T) {
	/* Normal case */
	actual, err := Pow(5, uint(0))
	if err != nil {
		t.Errorf("expected %v but given %v", nil, err)
	}
	if actual != 1 {
		t.Errorf("expected %v but given %v", 1, actual)
	}

	actual, err = Pow(6, uint(1))
	if err != nil {
		t.Errorf("expected %v but given %v", nil, err)
	}
	if actual != 6 {
		t.Errorf("expected %v but given %v", 6, actual)
	}

	actual, err = Pow(2, uint(4))
	if err != nil {
		t.Errorf("expected %v but given %v", nil, err)
	}
	if actual != 16 {
		t.Errorf("expected %v but given %v", 16, actual)
	}

	actual, err = Pow(-3, uint(3))
	if err != nil {
		t.Errorf("expected %v but given %v", nil, err)
	}
	if actual != -27 {
		t.Errorf("expected %v but given %v", -27, actual)
	}

	actual, err = Pow(13, uint(9))
	if err != nil {
		t.Errorf("expected %v but given %v", nil, err)
	}
	if actual != 10604499373 {
		t.Errorf("expected %v but given %v", 10604499373, actual)
	}
}

func TestPow_Overflow(t *testing.T) {
	/* Overflow case */
	_, err := Pow[int16, uint](2, 32)
	if err == nil {
		t.Errorf("expected error but given %v", nil)
	}

	_, err = Pow[uint8, uint](15, 3)
	if err == nil {
		t.Errorf("expected error but given %v", nil)
	}
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

func TestReversed(t *testing.T) {
	/* Length: odd or even */
	expected := []int{5, 4, 3, 2, 1}
	base := []int{1, 2, 3, 4, 5}
	base_copy := []int{1, 2, 3, 4, 5}
	actual := *Reversed(&base)
	if !reflect.DeepEqual(base, base_copy) {
		t.Errorf("expected %v but given %v", base_copy, base)
	}
	for i := 0; i < len(actual); i++ {
		if *actual[i] != expected[i] {
			t.Errorf("expected %v but given %v", expected, actual)
		}
	}

	expected = []int{9, 8, 7, 6}
	base = []int{6, 7, 8, 9}
	base_copy = []int{6, 7, 8, 9}
	actual = *Reversed(&base)
	if !reflect.DeepEqual(base, base_copy) {
		t.Errorf("expected %v but given %v", base_copy, base)
	}
	for i := 0; i < len(actual); i++ {
		if *actual[i] != expected[i] {
			t.Errorf("expected %v but given %v", expected, actual)
		}
	}

	/* non-ascii charactors */
	expected_rune := []rune("参弐壱")
	base_rune := []rune("壱弐参")
	base_rune_copy := []rune("壱弐参")
	actual_rune := *Reversed(&base_rune)
	if !reflect.DeepEqual(base_rune, base_rune_copy) {
		t.Errorf("expected %v but given %v", base_rune_copy, base_rune)
	}
	for i := 0; i < len(actual_rune); i++ {
		if *actual_rune[i] != expected_rune[i] {
			t.Errorf("expected %v but given %v", expected_rune, actual_rune)
		}
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
