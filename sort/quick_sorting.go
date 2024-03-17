// 快速排序
// 把数组的元素找一个基准元素，遍历一次把所有元素分为大于基准元素和小于基准元素的2份
// 对于分开的元素，重复上面的操作，直到所有的元素排序好(递归)
package main

import "fmt"

func sort(arr []int, low, high int) int {
  num := arr[high]

  i:= low - 1

  for j := low; j<high; j++ {

    if arr[j] <= num {
      i++
      arr[i], arr[j] = arr[j], arr[i]
    }
  }
  arr[i+1], arr[high] = arr[high], arr[i+1]
  return i+1
}

func QuickSort(arr []int, low, high int) {
  if low < high {
    num := sort(arr, low, high)
    QuickSort(arr, low, num -1)
    QuickSort(arr, num+1, high)
  }
}

func quickSort(arr []int) {
  QuickSort(arr,0, len(arr)-1)

}

func main() {
  arr := []int{5, 4, 2, 20, 8, 1, 44}
  quickSort(arr)
  fmt.Println(arr)
}
