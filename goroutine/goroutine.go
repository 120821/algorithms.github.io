// 协程 coroutine
// 轻量级的线程
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解释器/虚拟机层面的多任务
// 多个协程可能在一个或多个线程上运行

package main

import (
  "fmt"
  "time"
  "runtime"
)

func main() {
  var a [10]int
  for i:=0;i<10;i++ {
    go func(i int) {
      a[i]++
      runtime.Gosched()
      //fmt.Println(i)
      // printf是一个IO操作
    }(i)
  }
  //time.Sleep(time.Millisecond)
  time.Sleep(time.Minute)
  fmt.Println(a)
}

// 其他语言的协程
// c++ Boost.Coroutine
// java 不支持（第三方里可能支持）
// Python: 使用yield关键字实现协程
// Python3.5加入了async def 对协程原生支持

// 开启了一个进程
// 里边可以有很多的线程
// 线程里可能有很多的协程
// 调度器来管理每个线程有几个协程

// goroutine的定义
// 任何函数使用go关键字调用就能给调度器运行就成为一个协程
// 加上go的话，就不需要在定义的时候区分是否是异步函数（这是相对与python来讲的）
// 调度器在合适的点进行切换
// 使用-race 来检测数据访问冲突


// goroutine可能切换的点
// IO, select
// channel
// 等待锁
// 函数调用（有时候）
// runtime.Goshed()
// 只是可能会切换，不保证一定切换，不保证在其他点不切换
