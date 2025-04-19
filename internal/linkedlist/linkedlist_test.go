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

func TestRemoveFirst(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
		error    bool
	}

	testCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{2, 1},
			error:    false,
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
		error    bool
	}

	testCases := []testCase{
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 2},
			error:    false,
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
