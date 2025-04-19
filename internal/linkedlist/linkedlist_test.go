package linkedlist_test

import (
	"testing"

	"github.com/moh682/algorithms-go/internal/linkedlist"
)

func TestNewLinkedList(t *testing.T) {
	list := linkedlist.New[int]()
	if list.Len() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Len())
	}
}

func TestAdd(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
		error    bool
	}

	testCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
			error:    false,
		},
	}

	for _, tc := range testCases {
		list := linkedlist.New[int]()
		for _, v := range tc.input {
			list.Add(v)
		}

		if list.Len() != len(tc.expected) {
			t.Errorf("Expected list size to be %d, got %d", len(tc.expected), list.Len())
		}

		for _, v := range tc.expected {
			if !list.Contains(v) {
				t.Errorf("Expected list to contain %d, but it doesn't", v)
			}
		}

	}
}

func TestRemove(t *testing.T) {
	type testCase struct {
		input       []int
		removeValue int
		expected    []int
	}

	testCases := []testCase{
		{
			input:       []int{1, 2, 3},
			removeValue: 2,
			expected:    []int{3, 1},
		},
		{
			input:       []int{1, 2, 3},
			removeValue: 10,
			expected:    []int{3, 2, 1},
		},
		{
			input:       []int{},
			removeValue: 10,
			expected:    []int{},
		},
	}

	for _, tc := range testCases {
		list := linkedlist.New[int]()
		for _, v := range tc.input {
			list.Add(v)
		}

		list.Remove(tc.removeValue)

		if list.Len() != len(tc.expected) {
			t.Errorf("Expected list size to be %d, got %d", len(tc.expected), list.Len())
		}

		for _, v := range tc.expected {
			if !list.Contains(v) {
				t.Errorf("Expected list to contain %d, but it doesn't", v)
			}
		}

		if list.Contains(tc.removeValue) {
			t.Errorf("Expected list to not contain %d, but it does", tc.removeValue)
		}

	}
}

func TestRemoveFirst(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
	}

	testCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{2, 1},
		},
	}

	for _, tc := range testCases {
		list := linkedlist.New[int]()
		for _, v := range tc.input {
			list.Add(v)
		}

		list.RemoveFirst()

		if list.Len() != len(tc.expected) {
			t.Errorf("Expected list size to be %d, got %d", len(tc.expected), list.Len())
		}

		for _, v := range tc.expected {
			if !list.Contains(v) {
				t.Errorf("Expected list to contain %d, but it doesn't", v)
			}
		}

	}
}

func TestRemoveLast(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
	}

	testCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 2},
		},
	}

	for _, tc := range testCases {
		list := linkedlist.New[int]()
		for _, v := range tc.input {
			list.Add(v)
		}

		list.RemoveLast()

		if list.Len() != len(tc.expected) {
			t.Errorf("Expected list size to be %d, got %d", len(tc.expected), list.Len())
		}

		for _, v := range tc.expected {
			if !list.Contains(v) {
				t.Errorf("Expected list to contain %d, but it doesn't", v)
			}
		}

	}
}
