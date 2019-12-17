# 接口

接口类型是对其他类型行为的概括和抽象。通过使用接口，我们可以写出更加灵活和通用的函数，这些函数不用绑定在一个特定的类型实现上。

## 接口即约定

接口是一种抽象类型，它并没有暴露所包含数据的布局或者内部结构，当然也没有那些数据的基本操作，它所提供的仅仅是一些方法而已。

以 `fmt.Printf` 和 `fmt.Sprintf` 为例，前者把结果发到标准输出（标准输出其实是一个文件），后者把结果以 string 类型返回。

```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

func Printf(format string, args ...interface{}) (int, error) {
  return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
  var buf bytes.Buffer
  Fprintf(&buf, format, args...)
  return buf.string()
}
```

Fprintf 的前缀 F 指文件，表示格式化的输出会写入第一个实参所指代的文件。对于 Printf，第一个实参就是 os.Stdout，它属于 *os.File 类型。对于 Sprintf，尽管第一个实参不是文件，但它模拟了一个文件：&buf 就是一个指向内存缓冲区的指针，与文件类似，这个缓冲区也可以写入多个字节。

其实 Fprintf 的第一个形参也不是文件类型，而是 io.Writer 接口类型，其声明如下：

```go
package io

// Writer 接口封装了基础的写入方法
type Writer interface {
  // Write 从  p 向底层数据流写入 len(p) 个字节的数据
  // 返回实际写入的字节数（0 <= n <= len(p)）
  // 如果没有写完，那么会返回遇到的错误
  // 在 Write 返回 n < len(p) 时，err 必须为非 nil
  // Write 不允许修改 p 的数据，即使是临时修改
  //
  // 实现时不允许残留 p 的引用
  Write(p []byte) (n int, err error)
}
```

io.Writer 接口定义了 Fprintf 和调用者之间的约定。一方面，这个约定要求调用者提供的具体类型（比如 *os.File 或者 *bytes.Buffer）包含一个与其签名和行为一致的 Write 方法。另一方面，这个约定保证了 Fprintf 能使用任何满足 io.Writer 接口的参数。Fprintf 只需要能调用参数的 Write 函数，无须假设它写入的是一个文件还是一段内存。

因为 fmt.Fprintf 仅依赖于 io.Writer 接口所约定的方法，对参数的具体类型没有要求，所以我们可以用任何满足 io.Writer 接口的具体类型作为 fmt.Fprintf 的第一个实参。这种可以把一种类型替换为满足同一接口的另一种类型的特性称为可取代性，这也是面向对象语言的典型特征。