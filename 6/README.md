# 方法

## 方法声明

Go 和许多其他面向对象的语言不同，它可以将方法绑定到任何类型上。可以很方便地位简单的类型（如数字、字符串、slice、map，甚至函数等）定义附加的行为。

## 指针接收者的方法

由于主调函数会复制每一个实参变量，因此我们必须使用指针来传递变量的地址。

```go
func (p *Point) ScaleBy(factor float64) {
  p.X *= factor
  p.Y *= factor
}
```

这个方法的的名字是 (*Point).ScaleBy。圆括号是必需的；没有圆括号，表达式会被解析为 *(Point.ScaleBy)。

但是如果实参接收者是 *Point 类型，以 Point.Distance 方式调用 Point 类型的方法是合法的，编译器会自动插入一个隐式的 * 操作符。下面两个函数的调用效果是一样的：

```go
pptr.Distance(q)
(*pptr).Distance(q)
```

## 通过结构体内嵌组成类型

下面是个很好的示例。这个例子展示了简单的缓存实现，其中实现了两个包级别的变量——互斥锁和 map，互斥锁会保护 map 的数据。

```go
var (
  mu sync.Mutex
  mapping = make(map[string]string)
)

func Lookup(key string) string {
  mu.Lock()
  v := mapping[key]
  mu.Unlock()
  return v
}
```

下面这个版本的功能和上面完全相同，但是将两个相关变量放到了一个包级别的变量 cache 中：

```go
var cache = struct {
  sync.Mutex
  mapping map[string]string
} {
  mapping: make(map[string]string)
}

func Lookup(key string) string {
  cache.Lock()
  v := cache.mapping[key]
  cache.UnLock()
  return v
}
```

新的变量名更加贴切，而且 sync.Mutex 是内嵌的，它的 Lock 和 Unlock 方法也包含进了结构体中，允许我们直接使用 cache 变量本身进行加锁。

## 封装

如果变量或者方法是不能通过对象访问到的，这称作封装的变量或者方法。

Go 语言只有一种方式控制命名的可见性：定义的时候，首字母大写的标识符是可以从包中导出的，而首字母没有大写的则不导出。

Go 语言中的封装的单元是包而不是类型。无论是在函数内的代码还是方法内的代码，结构体类型内的字段对于同一个包中的所有代码都是可见的。

封装提供了三个优点：
- 第一，因为使用方不能直接修改对象的变量，所以不需要更多的语句来检查变量的值；
- 第二，隐藏实现细节可以防止使用方依赖的属性发生改变，使得设计者可以更加灵活地改变 API 的实现而不破坏兼容性；
- 第三，可以防止使用者肆意地改变对象内的变量；