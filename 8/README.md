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

- 你并不需要在结束时关闭所有的 channel，只需要在告知 goroutine 数据已经全部发送完成时关闭 channel 即可；
- channel 有专门的类型， <-chan int 代表一个只能发送不能接收的 channel 类型，chan<- int 代表一个只能接收不能发送的 channel 类型；

## Buffered Channels

- BufferedChannels 一般用于线程调度，处理几个线程之间的通信与平衡问题；下面这个例子将返回最先抵达的响应结果：

```go
func mirroredQuery() string {
  responses := make(chan string, 3)
  go func() { responses <- request("asia.gopl.io") }()
  go func() { responses <- request("europe.gopl.io") }()
  go func () { responses <- request("maericas.gopl.io") }()
  return <- responses
}
func request(hostname string) (response string) {  }
```

- 在上面这个例子中，如果使用的是 unbuffered channel，那么另外两个较慢的 goroutines 会一直尝试发送给一个没有接收者处理的 channel 中，垃圾回收机制将无法识别，会造成内存泄露！
- unbuffered 方便在 goroutine 中进行同步操作，而 buffered 适合在已知容量的情况下使用，这样不容易引发死锁和内存泄露问题。

## 并行循环

- 版本一：利用 channels

```go
func runTask(taskId int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(taskId)
}

func main() {
	task := make(chan int)
	for i := 0; i < 100; i++ {
		go func(taskId int) {
			runTask(taskId)
			task <- taskId
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-task
	}
}
```

- 版本二：sync.WaitGroup

```go
func makeThumbnails6(filenames <-chan string) int64 {
  sizes := make(chan int64)
  var wg sync.WaitGroup
  for f := range filenames {
    wg.Add(1)
    // worker
    go func(f string) {
      defer wg.Done()
      thumb, err := thumbnail.ImageFile(f)
      if err != nil {
        log.Println(err)
        return
      }
      info, _ := os.Stat(thumb)
      sizes <- info.Size()
    }(f)
  }
  // closer
  go func() {
    wg.Wait()
    close(sizes)
  }()
  var total int64
  for size := range sizes {
    total += size
  }
  return total
}
```

## 并发爬虫案例

```go
func main() {
  worklist := make(chan []string)

	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
```

- 由于 worklist 被初始化为了 channel 类型，所以对 command 参数的收集工作也需要运行在一个独立的进程中来避免死锁问题。
- 上面这个案例已经实现了高并发，但是会出现一个问题：
- `lookup blog.golang.org: no such host`：
  - 项目一次性创建了过多的网络连接，超过了系统限制数（单个文件可打开的进程数），从而引起了类似 DNS 查询出错的问题；
  - 解决方案是根据资源来限制并发数，比如把并发数限制在 20 以下；
  - 我们可以使用 buffered channel 实现对并发数的限制；
- 解决方案如下：

```go
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[:1] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
```

## Multiplexing with select

- select 可以同时接收多个 channel 的值；
- 如果我们希望在一个还没有准备好的 channel 中接收值并且不想发生堵塞时，我们可以使用 select 来避免堵塞；

## 并行遍历文件夹

- 实现类似于 Unix 系统的 `du` 命令功能；267