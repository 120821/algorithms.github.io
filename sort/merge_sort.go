// 归并排序
//  原始数组分成较小的数组，直到每个小数组只有一个位置，接着将小数组合并成较大的数组，最终合并为排序好的数组。
//复杂度：
//    时间复杂度：O(nlogn)O(nlogn)。
//    空间复杂度：O(n)O(n)。
package main

import "fmt"

func merge(left, right []int) (result []int){
  l, r := 0, 0
  for l < len(left) && r < len(right) {
    if left[l] < right[r] {
      result = append(result, left[l])
      l++
    } else {
      result = append(result, right[r])
      r++
    }
  }
  result = append(result, left[l:]...)
  result = append(result, right[r:]...)
  return


}

func mergeSort(arr []int) []int {
  if len(arr) <= 1{
    return arr
  }

  mid := len(arr) / 2
  left := mergeSort(arr[:mid])
  right:= mergeSort(arr[mid:])
  return merge(left, right)

}

func main() {
  arr := []int{4, 2, 5, 7, 29, 99, 3, 22, 25}
  result := mergeSort(arr)
  fmt.Println(result)

}
