package main

import "fmt"

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	if n == 0 {
// 		return
// 	}
// 	if m == 0 {
// 		for i := 0; i < n; i++ {
// 			nums1[i] = nums2[i]
// 		}
// 		return
// 	}
// 	tmpslice := make([]int, m+n)
// 	for i := 0; i < m; i++ {
// 		tmpslice[i+n] = nums1[i]
// 	}
// 	idx := 0
// 	i1, i2 := n, 0
// 	for i1 < m+n && i2 < n {
// 		if tmpslice[i1] > nums2[i2] {
// 			nums1[idx] = nums2[i2]
// 			i2++
// 		} else {
// 			nums1[idx] = tmpslice[i1]
// 			i1++
// 		}
// 		idx++
// 	}
// 	for i2 < n {
// 		nums1[idx] = nums2[i2]
// 		i2++
// 		idx++
// 	}
// 	for i1 < n+m {
// 		nums1[idx] = tmpslice[i1]
// 		i1++
// 		idx++
// 	}
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	sorted := make([]int, 0, m+n)
// 	p1, p2 := 0, 0

// 	for {
// 		if p1 == m {
// 			sorted = append(sorted, nums2[p2:]...)
// 			break
// 		}
// 		if p2 == n {
// 			sorted = append(sorted, nums1[p1:]...)
// 			break
// 		}
// 		if nums1[p1] < nums2[p2] {
// 			sorted = append(sorted, nums1[p1])
// 			p1++
// 		} else {
// 			sorted = append(sorted, nums2[p2])
// 			p2++
// 		}
// 	}
// 	copy(nums1, sorted)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main() {
	// a1 := []int{4, 0, 0, 0, 0, 0}
	// a2 := []int{1, 2, 3, 5, 6}
	a1 := []int{1, 2, 4, 5, 6, 0}
	a2 := []int{3}
	// a1 := make([]int, 1)
	// a2 := make([]int, 1)
	//a2[0] = 1
	merge(a1, 5, a2, 1)
	fmt.Println(a1)

}
