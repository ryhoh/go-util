package util

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	set := NewSet[int]()
	if set == nil {
		t.Errorf("expected Set[int] but given %v", nil)
	}
}

func TestAdd(t *testing.T) {
	set := NewSet[int]()
	expected := Set[int]{}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}

	set.Add(1)
	expected = Set[int]{1: struct{}{}}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}

	set.Add(2)
	expected = Set[int]{1: struct{}{}, 2: struct{}{}}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}
}

func TestAddAll(t *testing.T) {
	set := NewSet[int]()
	expected := Set[int]{1: struct{}{}, 2: struct{}{}}

	set.AddAll(&[]int{1, 2})
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}
}

func TestSetFromSlice(t *testing.T) {
	set := *SetFromSlice(&[]int{1, 2})
	expected := Set[int]{1: struct{}{}, 2: struct{}{}}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}
}

func TestRemove(t *testing.T) {
	set := *SetFromSlice(&[]int{1, 2})
	expected := Set[int]{1: struct{}{}, 2: struct{}{}}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}

	set.Remove(2)
	expected = Set[int]{1: struct{}{}}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}

	set.Remove(1)
	expected = Set[int]{}
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}
}

func TestRemoveAll(t *testing.T) {
	set := *SetFromSlice(&[]int{1, 2})
	expected := Set[int]{}

	set.RemoveAll(&[]int{1, 2})
	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v but given %v", expected, set)
	}
}

func TestContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	if !set.Contains(1) {
		t.Errorf("expected %v but given %v", true, false)
	}
	if set.Contains(4) {
		t.Errorf("expected %v but given %v", false, true)
	}
}

func TestString(t *testing.T) {
	set := NewSet[int]()
	expected := "Set[]"
	actual := set.String()
	if actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	set.Add(1)
	expected = "Set[1]"
	actual = set.String()
	if actual != expected {
		t.Errorf("expected %v but given %v", expected, actual)
	}

	set.Add(2)
	set.Add(3)
	set.Add(4)
	expected_list := []string{"1", "2", "3", "4"}
	actual_list := strings.Split(
		strings.Replace(
			strings.Replace(set.String(), "Set[", "", 1), "]", "", 1),
		", ",
	)
	sort.Slice(actual_list, func(a, b int) bool { return actual_list[a] < actual_list[b] })
	if !reflect.DeepEqual(actual_list, expected_list) {
		t.Errorf("expected %v but given %v", expected_list, actual_list)
	}
}
