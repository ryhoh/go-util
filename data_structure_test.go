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

func TestString_Set(t *testing.T) {
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

func TestPush(t *testing.T) {
	slice := []int{}
	expected := []int{1}
	Push(&slice, 1)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}

	Push(&slice, 2)
	expected = []int{1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
}

func TestPop(t *testing.T) {
	slice := []int{1, 2}
	expected := []int{1}
	popped := *Pop(&slice)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if popped != 2 {
		t.Errorf("expected %v but given %v", 2, popped)
	}

	popped = *Pop(&slice)
	expected = []int{}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if popped != 1 {
		t.Errorf("expected %v but given %v", 2, popped)
	}

	popped_p := Pop(&slice)
	expected = []int{}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if popped_p != nil {
		t.Errorf("expected %v but given %v", nil, popped_p)
	}
}

func TestTop(t *testing.T) {
	slice := []int{1, 2}
	expected := 2
	top := *Top(&slice)
	if top != expected {
		t.Errorf("expected %v but given %v", expected, top)
	}

	slice = []int{}
	top_p := Top(&slice)
	if top_p != nil {
		t.Errorf("expected %v but given %v", nil, top)
	}
}

func TestEnqueue(t *testing.T) {
	slice := []int{}
	expected := []int{1}
	Enqueue(&slice, 1)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}

	Enqueue(&slice, 2)
	expected = []int{1, 2}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
}

func TestDequeue(t *testing.T) {
	slice := []int{2, 1}
	expected := []int{1}
	dequeued := *Dequeue(&slice)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if dequeued != 2 {
		t.Errorf("expected %v but given %v", 2, dequeued)
	}

	dequeued = *Dequeue(&slice)
	expected = []int{}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if dequeued != 1 {
		t.Errorf("expected %v but given %v", 2, dequeued)
	}

	dequeued_p := Dequeue(&slice)
	expected = []int{}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v but given %v", expected, slice)
	}
	if dequeued_p != nil {
		t.Errorf("expected %v but given %v", nil, dequeued_p)
	}
}

func TestFront(t *testing.T) {
	slice := []int{2, 1}
	expected := 2
	front := *Front(&slice)
	if front != expected {
		t.Errorf("expected %v but given %v", expected, front)
	}

	slice = []int{}
	front_p := Front(&slice)
	if front_p != nil {
		t.Errorf("expected %v but given %v", nil, front)
	}
}
