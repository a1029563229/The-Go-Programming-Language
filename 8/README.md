# Goroutines and Channels

## Goroutines

- 在 Go 中，每个并行的线程被称为 goroutines；
- 一个 Goroutines 可以简单想象成系统中的一个线程。
- 当一个项目启动时，只有一个 goruntine，那就是主线程（main function），你可以用 `go` 关键字来启动一个新的线程，就像下面这个例子：

```go
f() // 执行函数，等待执行结果的返回
go f() // 启动新的进程执行函数，不会等待执行结果的返回
```

### Example: Concurrent Clock Server
### Example：Concurrent Echo Server

## Channels

- 无缓冲区 channel
  - 当某个 goroutines 的 channel 发送了一个值后，如果没有对应的接收操作，那这个 goroutines 会处于 pending 状态，直到有对应的 channel 接收了发送的值；反之也是一样，如果在某个 goroutines 进行了一个 channel 的接收操作后，如果没有接收到对应 channel 发送的值，这个 goroutines 也会处于 Pending 状态，直到接收到 channel 的值；
  - 使用无缓冲区 channel 进行通信可以使得在 goroutines 的发送和接收操作变成同步，所以无缓冲区 channels 有时候也叫做同步 channels。

- Pipelines
  - channels 也可以用将输出变为输入的形式连接几个 goroutines，这种模式称作 Pipelines。

![pipelines](http://shadows-mall.oss-cn-shenzhen.aliyuncs.com/images/blogs/other/Jietu20191206-173550.png)

- 使用 close 关闭 channel 时，所有的发送操作都会引起 panic，但是接收操作会接收到一个对应类型的零值。
- 检测 channel 是否关闭的方法：
```go
go func() {
  for {
    x, ok := <-naturals
    if !ok {
      break // channel 已经被关闭
    } 
    squares <- x * x
  }
  close(squares)
}()
```