package main
import "sort"
import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 将 nums2 追加到 nums1 的末尾
	copy(nums1[m:], nums2)

	// 对 nums1 进行排序
	sort.Ints(nums1)
}

func main() {
  m := 4
  nums1 := make([]int, m)
  copy(nums1, []int{5, 6, 28, 29})
  nums2 := []int{18, 19, 23, 27}
  n := len(nums2)
  merge(nums1, m, nums2, n)
  fmt.Println(nums1)
}
