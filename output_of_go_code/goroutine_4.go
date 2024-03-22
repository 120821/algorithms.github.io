package main

import "fmt"
import "time"

func main() {
  strs := []string{"one", "two", "three"}
  //strs := []string{"one", "two", "three", "four"}
  // 由于闭包的特性，所有的goroutines共享相同的变量s。
  // 当goroutines开始执行时，它们引用的s变量已经完成了迭代，此时s的值为"three"。
  // 在循环的最后一次迭代时，所有的goroutines都已经创建，在执行fmt.Printf("%s ", s)时都引用的是相同的s变量。
  // 由于goroutines是并发执行的，它们在开始执行时s的值已经是"three"，因此每个goroutine都会打印"three"。
  // 所以，增加了"four"后每次输出的都是"four"
  for _, s := range strs {
    go func() {
      fmt.Printf("%s ", s)
    }()
  }
  time.Sleep(3 * time.Second)
}
// three three three 等待3秒
// four four four four 等待3秒
