package main
// 快速排序
import "fmt"
// 基准
// 小大，大小
// 左边+ 右边的
// 递归
func sort(arr []int) []int{
  if len(arr) <= 1 {
    return arr
  }
  //length := len(arr)
  result := []int{}
  num := arr[0] // 4
  left, right := []int{}, []int{}
  mid := 0
  for i := range arr {
    if arr[i] < num {
      right = append(right, arr[i])
    } else if arr[i] == num {
      mid = num
    } else {
      left = append(left, arr[i])
    }
  }

  result = append(result, sort(left)...)
  result = append(result, mid)
  result = append(result, sort(right)...)

  // left 4 2 1
  return  result
}

func main() {
  arr := []int{4, 2, 5, 1, 8}
  result := sort(arr)
  fmt.Println(result)
}
