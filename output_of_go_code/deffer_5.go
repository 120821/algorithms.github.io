// defer语句的执行顺序是后进先出的，也称为"栈"的特性
// defer语句的执行顺序与其定义顺序无关，而是与其所在的函数返回时的执行顺序有关。
// 一个函数中存在多个defer语句时，会被压入一个栈中，在函数返回时按照后进先出的顺序执行。
package main

import "fmt"

func increaseA() int {
  var i int // 0
  // defer会在函数返回之前被延迟执行
  defer func() {
    i++ // 1
  }()
  // increaseA函数在defer语句执行之前就已经返回了
  return i
}

func increaseB() (r int) {
  // r是函数返回值的一部分，初始值也是0
  // defer会在函数返回之前被延迟执行
  // 匿名函数使变量r的值加1
  // r是作为返回值命名的，所以在函数返回时，它的值被保留下来
  defer func() {
    r++ // 1
  }()
  return r
}

func main() {
  fmt.Println(increaseA()) // 0
  fmt.Println(increaseB()) // 1
}

// 0
// 1

// increaseB函数中的defer语句会在函数返回之前被延迟执行。
// r的初始值为0，函数返回之前会被自增为1。会输出1。

// increaseA函数中的defer语句会在函数返回之前被延迟执行。
// i是在匿名函数执行之后返回的，所以初始值0会被返回。会输出0。
// 而且，无论是先调用increaseA increaseB都不影响increaseA increaseB函数的返回值。
