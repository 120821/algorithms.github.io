package main

import (
  "fmt"
  "math"
)

func reverseInt(x int) int {
  sign := 1
  if x < 0 {
    sign = -1
    x = -x
  }
  // 反转整数
  reversed := 0
  for x > 0 {
    reversed = reversed *10 + x %10
    x /= 10
  }
  reversed *= sign
  if reversed < math.MinInt32 || reversed > math.MaxInt32 {
    return 0
  }
  return reversed
}

func main() {
  x := -1234
  y := 1234
  reversed_x := reverseInt(x)
  reversed_y := reverseInt(y)
  fmt.Println(reversed_x)
  fmt.Println(reversed_y)

}
