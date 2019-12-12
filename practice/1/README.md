# 入门 - 练习题

- [1.1：修改 echo 程序输出 os.Args[0]，即命令的名称。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.1)
- [1.2：修改 echo 程序，输出参数的索引和值，每行一个。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.2)
- [1.3：尝试测量可能低效的程序和使用 strings.Join 的程序在执行时间上的差异。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.3)
- [1.4：修改 dup2 程序，输出出现重复行的文件的名称。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.4)
- [1.7：函数 io.Copy(dst, src) 从 src 读，并且写入 dst。使用它代替 ioutil.ReadAll 来复制响应内容到 os.Stdout，这样不需要装下整个响应数据流的缓冲区。确保检查 io.Copy 返回的错误结果。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.7)
- [1.8：修改 fetch 程序添加一个 http:// 前缀（假如该 URL 参数缺失协议前缀）。可能会用到 strings.HasPrefix。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.8)
- [1.9：修改 fetch 来输出 HTTP 的状态码，可以在 resp.Status 中找到它。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.9)
- [1.10：找一个产生大量数据的网站。连续两次运行 fetchall，看报告的时间是否会有大的变化，调查缓存情况。每一次获取的内容一样吗？修改 fetchall 将内容输出到文件，这样可以检查它是否一致。](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.10)
- [1.11：使用更长的参数列表来尝试 fetchall，例如使用 alexa.com 排名前 100 万的网站。如果一个网站没有响应，程序的行为是怎样的？](https://github.com/a1029563229/The-Go-Programming-Language/tree/master/practice/1/1.11)