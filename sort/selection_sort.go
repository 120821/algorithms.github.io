package main

import "fmt"

func selectionSort (arr []int) []int{
  length := len(arr)
  for i:=0;i<length-1; i++ {
    min_index := i
    for j:=i+1;j<length;j++{
      if arr[j] < arr[min_index] {
        min_index = j
      }
      if min_index != i {
        arr[i], arr[min_index] = arr[min_index], arr[i]
      }
    }
  }
  return arr
}

func main() {
  arr := []int{4, 2, 21, 12, 29, 24}
  selectionSort(arr)
  fmt.Println(arr)
}
