package view

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

//有序
func SortedList1(left, right *Node) *Node {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = SortedList1(left.Next, right)
		return left
	} else {
		right.Next = SortedList1(left, right.Next)
		return right
	}

}

func (n *Node) printNode() {
	for n != nil {
		fmt.Printf("%d \t ", n.Val)
		n = n.Next
	}
	fmt.Println()
}

func SortedList(nums1 []int, m int, nums2 []int, n int) {
	// 将最大元素移到最末端
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	// 处理没有处理完的元素 m > 0 nums1 无需处理
	// n > 0  nums2[:n] 迁移至nums1
	if n > 0 {
		nums1 = append(nums1[:0], nums2[:n]...)
	}
}
