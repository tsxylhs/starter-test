package view

import (
	"testing"
)

func TestCombineSortedList(t *testing.T) {
	n1 := &Node{
		Val: 9,
		Next: &Node{
			Val: 11,
			Next: &Node{
				Val: 12,
			},
		},
	}
	n2 := &Node{
		Val: 13,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 18,
			},
		},
	}
	result := SortedList1(n1, n2)
	result.printNode()

}
