## 测试

`go test` 子命令是 Go 语言包的测试驱动程序，这些包根据某些约定组织在一起。`go test` 会编译以 `_test.go` 结尾的文件。

在 `*_test.go` 文件中，三种函数需要特殊对待，即功能测试函数、基准测试函数和示例函数。功能测试函数是以 `Test` 前缀命名的函数，用来检测一些程序逻辑的正确性。基准测试函数的名称以 `Benchmark` 开头，用来测试某些操作的性能，`go test` 汇报操作的平均执行时间。示例函数的名称，以 `Example` 开头，用来提供机器检查过的文档。

`go test` 工具扫描 `*_test.go` 文件来寻找特殊函数，并生成一个临时的 `main` 包来调用它们，然后编译和运行，并汇报结果，最后清空临时文件。

## 白盒测试与黑盒测试

测试的分类之一是基于对所要进行测试的包的内部了解程度。黑盒测试假设测试者对包的了解仅通过公开的 API 和文档，而包的内部逻辑则是不透明的。相反，白盒测试可以访问包的内部函数和数据结构，并且可以做一些常规用户无法做到的观察和改动。

这两种方法是互补的。黑盒测试通常更加健壮，每次程序更新后基本不需要修改。它们也会帮助测试的作者关注包的用户并且能够发现 API 设计的缺陷。反之，白盒测试可以对实现的特定之处提供更详细的覆盖测试。