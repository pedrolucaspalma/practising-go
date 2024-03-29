package linkedlist

import (
	"testing"
)

var testIntVals = []int{1, 3, 5, 6, 12, 7, 14, 99}
var testStringVals = []string{"lala1", "lala2", "lala3", "Foo", "Bar", "Baz"}

func TestCreate(t *testing.T) {
	for idx, val := range testIntVals {
		root := Create[int](val)

		if root.next != nil {
			t.Errorf("Val of idx %d in testIntVals wasnt created as root value. Expected nil, got %d \n", idx, root.next)
		}

		if root.Val != val {
			t.Errorf("Val of idx %d in testIntVals wasnt created with proper value. Expected %d, got %d \n", idx, val, root.Val)
		}
	}
	for idx, val := range testStringVals {
		root := Create[string](val)

		if root.next != nil {
			t.Errorf("Val of idx %v in testIntVals wasnt created as root value. Expected nil, got %v \n", idx, root.next)
		}

		if root.Val != val {
			t.Errorf("Val of idx %d in testIntVals wasnt created with proper value. Expected %v, got %v \n", idx, val, root.Val)
		}
	}
}

func TestPushIntegers(t *testing.T) {
	var root = Create[int](testIntVals[0])

	for idx, val := range testIntVals {
		if idx == 0 {
			continue
		}
		root.Push(val)
	}

	aux := *root
	for idx, val := range testIntVals {
		if aux.Val != val {
			t.Errorf("Val of idx %d in testIntVals wasnt created with proper value. Expected %v, got %v \n", idx, val, aux.Val)
		}
		if idx != len(testIntVals)-1 {
			aux = *aux.next
		}
	}
}

func TestPopIntegers(t *testing.T) {
	var root = Create[int](testIntVals[0])

	for idx, val := range testIntVals {
		if idx == 0 {
			continue
		}
		root.Push(val)
	}

	aux := *root
	for idx, val := range testIntVals {
		if aux.Val != val {
			t.Errorf("Val of idx %d in testIntVals wasnt created with proper value. Expected %v, got %v \n", idx, val, aux.Val)
		}
		if idx != len(testIntVals)-1 {
			aux = *aux.next
		}
	}
}
