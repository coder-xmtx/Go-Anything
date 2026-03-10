# 03-流程控制

## if else

和其他语言类似

```go
score := 80
if score >= 80 {
    fmt.Println("优秀")
} else if score >= 60 {
    fmt.Println("良好")
} else {
    fmt.Println("不及格")
}
```

Go语言的if条件可以包含一个简单语句（如变量声明），常用于错误处理：

```go
if err := process(); err != nil {
    fmt.Println("错误:", err)
}
```

## for range 循环

和其他语言差不多

```go
orders := []string{"A", "B", "C"}
for index, value := range orders {
    fmt.Println("处理的订单是:", value, "订单序号是:", index)
}
```

如果不需要索引，可以用下划线`_`忽略

```go
for _, value := range orders {
    fmt.Println("处理的订单:", value)
}
```

## for 循环作为 while 使用

Go语言没有直接的while关键字，可以用for循环来代替

```go
i := 0
for i < 5 {
    i++
    fmt.Println("处理的订单是:", i)
}
```

## switch值匹配

Go语言的case支持多值，用逗号分隔

```go
statusCode := 200
switch statusCode {
case 200, 201:
    fmt.Println("OK")
case 404:
    fmt.Println("Not Found")
case 500:
    fmt.Println("Internal Server Error")
default:
    fmt.Println("Unknown")
}
```

==【重要特性】==

Go 的`switch`每个`case`默认带`break`，不会自动向下穿透。如果需要穿透，需显式使用`fallthrough`：

```go
switch num {
case 1:
    fmt.Println("1")
    fallthrough  // 会继续执行 case 2 的内容
case 2:
    fmt.Println("2")
}
```

## switch 无表达式

【例子】重构上面 if else 的例子

```go
score2 := 80
switch {
case score2 >= 80:
    fmt.Println("优秀")
case score2 >= 60:
    fmt.Println("良好")
default:
    fmt.Println("不及格")
}
```

## select 语句

这一块要和channel的知识一块学习，暂时跳过

```go
ch := make(chan string, 1)
go func() {
    time.Sleep(2 * time.Second)
    ch <- "finish"
}()
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}

// goroutine 需要 2 秒才能发送数据，
// 而 time.After(1 * time.Second) 会在 1 秒后产生一个事件，
// 所以 select 会执行 timeout 分支。
```

这里给出deepseek的讲解：

### 什么是 channel？

- **channel** 是 Go 中**goroutine 之间通信的管道**。你可以把它想象成一个队列：
  - 一个 goroutine 可以通过 `ch <- value` 向 channel **发送**数据。
  - 另一个 goroutine 可以通过 `<-ch` 从 channel **接收**数据。
- channel 是类型相关的，比如 `chan string` 只能传递字符串。
- 它有两种类型：
  - **无缓冲 channel**：发送和接收必须同时准备好，否则阻塞。
  - **有缓冲 channel**：如 `make(chan string, 1)` 创建了一个容量为 1 的缓冲，发送方可以在缓冲未满时立即发送，接收方在缓冲非空时立即接收。

### 为什么需要 `select`？

`select` 用于**同时等待多个 channel 操作**（发送或接收）。当多个 goroutine 可能在不同时间向不同 channel 发送数据时，`select` 可以让你优雅地处理最先到达的那个。

在你的例子中：

1. 你创建了一个带缓冲的 channel `ch`。
2. 启动一个 goroutine，它会在 2 秒后向 `ch` 发送 `"finish"`。
3. `select` 同时等待：
   - `case msg := <-ch`：等待从 `ch` 接收数据。
   - `case <-time.After(1 * time.Second)`：等待 1 秒超时事件。
4. 由于 goroutine 需要 2 秒，而超时是 1 秒，所以 1 秒后超时事件触发，`select` 执行超时分支，打印 `"timeout"`。

### `select` 的行为规则：

- 如果有多个 `case` 同时满足（例如多个 channel 都有数据），`select` 会**随机选择一个**执行。
- 如果没有任何 `case` 满足，且没有 `default` 分支，则 `select` 会**阻塞**，直到某个 `case` 可以执行。
- 如果有 `default` 分支，且所有 `case` 都不满足，则立即执行 `default`（非阻塞）。

### 学习建议

`select` 和 `channel` 是 Go 并发模型的核心，建议你先系统学习 goroutine 和 channel 的基础知识（比如通过 `go` 关键字启动并发任务，理解 channel 的阻塞机制），然后再回头研究 `select`，会更容易掌握。
