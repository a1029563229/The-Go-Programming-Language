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

## 实现接口

一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型定义中的所有方法。

如果一个类型实现了一个接口要求的所有方法，那么这个类型实现了这个接口。

空接口类型 `interface{}` 是不可缺少的，因为空接口类型对其实现类型没有任何要求，所以我们可以把任何值赋给空接口类型。

## 接口值

从概念上来讲，一个接口类型的值（简称接口值）其实有两个部分：一个具体类型和该类型的一个值。二者称为接口的动态类型和动态值。

在 Go 语言中，变量总是初始化为一个特定的值，接口也不例外。接口的零值就是把它的动态类型和值都设置为 `nil`。

接口值可以用 == 和 != 操作符来做比较。如果两个接口值都是 nil 或者两者的动态类型完全一致且两者动态值相等（使用动态类型的 == 操作符来做比较），那么两个接口值相等。因为接口值是可以比较的，所以它们可以作为 map 的键，也可以作为 switch 语句的操作数。

## http.Handler 接口

```go
// net/http
package http

type Handler interface {
  ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

`ListenAndServe` 函数需要一个服务器地址，比如 `localhost:8000`，以及一个 `Handler` 接口的实例（用来接受所有的请求）。这个函数会一直运行，直到服务出错（或者启动就失败了）时返回一个非空的错误。

`net/http` 包提供了一个请求多工转发器 `ServeMux`，用来简化 URL 和处理程序之间的关联。一个 ServeMux 把多个 `http.Handler` 组合成单个 `http.Handler`。

表达式 `http.HandlerFunc(db.list)` 其实是类型转换，而不是函数调用。注意，`http.HandlerFunc`  是一个类型，它有如下定义：

```go
// net/http
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```

`HandlerFunc` 演示了 Go 语言接口机制的一些不常见特性。它不仅是一个函数类型，还拥有自己的方法，也满足接口 `http.Handler`。它的 `ServeHTTP` 方法就调用函数本身，所以 `HandlerFunc` 就是一个让函数值满足接口的一个适配器。

为简便起见，`net/http` 包提供了一个全局的 `ServeMux` 实例 `DefaultServeMux`，以及包级别的注册函数 `http.Handle` 和 `http.HandleFunc`。要让 `DefaultServeMux` 作为服务器的主处理程序，无须把它传给 `ListenAndServe`，直接传 `nil` 即可。

```go
func main() {
  db := database{"shoes": 50, "socks": 5}
  http.HandleFunc("/list", db.list)
  http.HandleFunc("/price", db.price)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```

Web 服务器每次都用一个新的 `goroutine` 来调用处理程序，所以处理程序必须注意并发问题。

## error 接口

`error` 是一个接口类型，包含一个返回错误消息的方法：

```go
type error interface {
  Error() string
}
```

构造 `error` 最简单的方法是调用 `errors.New`，它会返回一个包含指定的错误消息的新 `error` 实例。完整的 `error` 包只有如下 4 行代码：

```go
package errors

func New(text string) error { return &errorString{text} }

type errorString struct { text string }

func (e *errorString) Error() string { return e.text }
```

直接调用 `errors.New` 比较罕见，因为有一个更易用的封装函数 `fmt.Errorf`，它还额外提供了字符串格式化功能：

```go
package fmt

import "errors"

func Errorf(format string, args ...interface{}) error {
  return errors.New(Sprintf(format, args...))
}
```

## 类型断言

类型断言是一个作用在接口值上的操作，写出来类似于 `x.(T)`，其中 `x` 是一个接口类型的表达式，而 `T` 是一个类型（称为断言类型）。类型断言会检查作为操作数的动态类型是否满足指定的断言类型。

```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File) // 成功：ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // 失败：!ok, b == nil
```

通过类型断言来识别错误：

```go
import (
  "errors"
  "syscall"
)
var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
  if pe, ok := err.(*PathError); ok {
    err = pe.Err
  }
  return err == syscall.ENOENT || err == ErrNotExist
}

// 实际调用
_, err := os.Open("/no/such/file")
fmt.Println(os..IsNotExist(err)) // "true"
```

## 通过接口类型断言来查询特性

```go
func writeHeader(w io.Writer, contentType string) error {
  if _, err := w.Write([]byte("Content-type: ")); err != nil {
    return err
  }
}
```

因为 `Write` 方法需要一个字节 `slice`，而我们想写入的是一个字符串，所以 `[]byte(...)` 转换就是必需的。这种转换需要进行内存分配和内存复制，但复制后的内存又会被马上抛弃。我们能否避开内存分配呢？

```go
// writeString 将 s 写入 w
// 如果 w 有 WriteString 方法，那么将直接调用该方法
func writeString(w io.Writer, s string) (n int, err error) {
  type stringWriter interface {
    WriteString(string) (n int, err error)
  }
  if sw, ok := w.(stringWriter); ok {
    return sw.WriteString(s) // 避免了内存复制
  }
  return w.Write([]byte(s)) // 分配了临时内存
}

func writeHeader(w io.Writer, contentType string) error {
  if _, err := writeString(w, "Content-type: "); err != nil {
    return err
  }
}
```

这个方法也用在了 `fmt.Printf` 中，用于从通用类型中识别出 `error` 或者 `fmt.Stringer`。在 `fmt.Fprintf` 内部，有一步是把单个操作数转换为一个字符串，如下所示：

```go
package fmt

func formatOneValue(x interface{}) string {
  if err, ok := x.(error); ok {
    return err.Error()
  }
  if str, ok := x.(Stringer); ok {
    return str.String()
  }
  // ... 所有其他类型
}
```

使用 `type` 关键字进行类型断言：

```go
func sqlQuote(x interface{}) string {
  switch x := x.(type) {
    case nil:
      return "NULL"
    case int, uint:
      return fmt.Sprintf("%d", x) // 这里 x 类型为 interface{}
    case bool:
      if x {
        return "TRUE"
      }
      return "FALSE"
    case string:
      return sqlQuoteString(x) // （未显示具体代码）
    default:
      panic(fmt.Sprintf("unexpected type %T: %v", x, x))
  }
}
```

## 一些建议

当设计一个新包时，一个新手 Go 程序员会首先创建一系列接口，然后再定义满足这些接口的具体类型。这种方式会产生很多接口，但这些接口只有一个单独的实现。不要这样做。这种接口是不必要的抽象，还有运行时的成本。可以用导出机制来限制一个类型的哪些方法或结构体的哪些字段是对包外可见的。仅在有两个或多个具体类型需要按统一的方法处理时才需要接口。

设计新类型时越小的接口越容易满足。一个不错的接口设计经验是仅要求你需要的。