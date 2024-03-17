package main

import "fmt"

func insertionSort(arr []int) {
  n := len(arr)
  for i:=0;i<n;i++{
    key := arr[i]
    j := i-1
    for j >= 0 && arr[j] > key {
      arr[j+1] = arr[j]
      j-=1
    }
    arr[j+1] = key
  }
}

func main() {
  arr := []int{5, 4, 2, 1, 7, 22, 21, 26}
  insertionSort(arr)
  fmt.Println(arr)
}
