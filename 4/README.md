# 复合数据类型

数据和结构体都是聚合类型，它们的值由内存中的一组变量构成。数组的元素具有相同的类型，而结构体中元素数据类型则可以不同。数组和结构体的长度都是固定的。反之，slice 和 map都是动态数据结构，它们的长度在元素添加到结构中时可以动态增长。

## 数组

- 数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列。由于数组的长度固定，所以在 Go 里面很少直接使用。slice 的长度可以增长和缩短，在很多场合下使用得更多。
- 如果一个数组的元素类型是可比较的，那么这个数组也是可比较的，这样我们就可以直接使用 == 操作符来比较两个数组，比较的结果是两边元素的值是否完全相同。
- 当调用一个函数的时候，每个传入的参数都会创建一个副本，然后复制给对应的函数变量，所以函数接收的是一个副本，而不是原始的参数。使用这种方式传递大的数组会变得很低效，并且在函数内部对数组的任何修改都仅影响副本，而不是原始数组。这种情况下，Go 把数组和其他的类型都看成值传递。而在其他的语言中，数组是隐式的使用引用传递。
- 当然，也可以显式地传递一个数组的指针给函数，这样在函数内部对数组的任何修改都会反映到原数组上面。
- 使用数组指针是高效的，同时允许被调用的函数修改调用方数组中的元素。

```go
func zero(ptr *[32]byte) {
  for i := range ptr {
    ptr[i] = 0
  }
}
```

## slice

- slice 表示一个拥有相同类型元素的可变长度的序列。slice 通常写成 []T，其中元素的类型都是 T；它看上去像没有长度的数组类型。
- Go 的内置函数 len 和 cap 用来返回 slice 的长度和容量。
- slice 操作符 s[i:j] （其中 0 <= i <= j <= cap(s)）创建了一个新的 slice，这个新的 slice 引用了序列 s 中从 i 到 j - 1 索引位置的所有元素。
- slice 无法做比较，标准库里面提供了高度优化的函数 bytes.Equal 来比较两个字节 slice（[]byte）；

- append 函数
- 内置函数 append 用来将元素追加到 slice 的后面，append 函数对理解 slice 的工作原理很重要。
- 每一次 append 调用都必须检查 slice 是否仍有足够容量来存储数组中的新元素。如果 slice 容量足够，那么它就会定义一个新的 slice（仍然引用原始底层数组），然后将新元素 y 复制到新的位置，并返回这个新的 slice。输入参数 slice x 和函数返回值 slice z 拥有相同的底层数组。
- 如果 slice 的容量不够容纳增长的元素，append 函数必须创建一个拥有足够容量的新的底层数组来存储新元素，然后将元素从 slice x 复制到这个数组，再将新元素 y 追加到数组后面。返回值 slice z 将和输入参数 slice x 引用不同的底层数组。
- 处于效率的考虑，新创建的数组容量会比实际容纳 slice x 和 slice y 所需要的最小长度更大一点。在每次数组容量扩展时，通过扩展一倍的容量来减少内存分配的次数，这样也可以保证追加一个元素所消耗的是固定时间。

```go
func appendInt(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    // slice 仍有增长空间，扩展 slice 内容
    z = x[:zlen]
  } else {
    // slice 已无空间，为它分配一个新的底层数组
    // 为了达到分摊线性复杂性，容量扩展一倍
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x) // 内置 copy 函数
  }
  z[len(x)] = y
  return z
}
```

## map

- 散列表是一个拥有键值对元素的无序集合。在这个集合中，键的值是唯一的。
- map 中所有的键都拥有相同的数据类型，同时所有的值也都拥有相同的数据类型。
- map 元素不是一个变量，不可以获取它的地址。我们无法获取 map 元素的地址的一个原因是 map 的增长可能会导致已有元素被重新散列到新的存储位置，这样就可能使得获取的地址无效；
- map 中元素的迭代顺序是不固定的；这个是有意为之的，这样可以使得程序在不同的散列算法实现下变得健壮；
- 和 slice 一样，map 不可比较，唯一合法的比较就是和 nil 做比较。为了判断两个 map 是否拥有相同的键和值，必须写一个循环：

```go
func equal(x, y map[string]int) bool {
  if len(x) != len(y) {
    return false
  }
  for k, xv := range x {
    if yv, ok := y[k]; !ok || yv != xv {
      return false
    }
  }
  return true
}
```

## 结构体

- 结构体是将零个或者多个任意类型的命名变量组合在一起的聚合数据类型。
- 命名结构体类型 S 不可以定义一个拥有相同结构体类型 S 的成员变量，但是 S 中可以定义一个 S 的指针类型，即 *S，这样我们就可以创建一些递归数据结构，比如链表和树。
- 没有任何成员变量的结构体称为空结构体，写作 struct{}。它没有长度，也不携带任何信息，但是有的时候会很有用。
- 结构体类型的值可以作为参数传递给函数或者作为函数的返回值，出于效率的考虑，大型的结构体通常都使用结构体指针的方式直接传递给函数或者从函数中返回。

```go
func Bonus(e *Employee, percent int) int {
  return e.Salary * percent / 100
}
```

- 由于通常结构体都通过指针的方式使用，因此可以使用一种简单的方式来创建、初始化一个 struct 类型的变量并获取它的地址：

```go
pp := &Point{1, 2}
```

- 结构体比较
- 如果结构体的所有成员变量都可以比较，那么这个结构体就是可比较的。两个结构体的比较可以使用 == 或者 !=。
- 和其他可比较的类型一样，可比较的结构体类型都可以作为 map 的键类型。

- 结构体嵌套和匿名成员

```go
// Go 允许我们定义不带名称的结构体成员，只需要指定类型即可；这种结构体成员称做匿名成员。这个结构体成员的类型必须是一个命名类型或者指向命名类型的指针
type Point {
  x, y int
}

type Circle struct {
  Point
  Radius int
}

type Wheel struct {
  Circle
  Spokes int
}

var w Wheel
w.X = 8 // 等价于 w.Circle.Point.X = 8
w.Y = 8 // 等价于 w.Circle.Point.Y = 8
w.Radius = 5 // 等价于 w.Circle.Radius = 5
w.Spokes = 20

// 结构体字面量必须遵循形状类型的定义，所以我们使用下面的两种形式来初始化：
w = Wheel{Circle{Point{8, 8}, 5}, 20}

w = Wheel{
  Circle: Circle{
    Point: Point{X: 8, Y: 8},
    Radius: 5,
  },
  Spokes: 20,
}
```

- 以快捷方式访问匿名变量的内部变量同样适用于访问匿名成员的内部方法。因此，外围的结构体类型获取的不仅是匿名成员的内部变量，还有相关的方法。这个机制就是从简单类型对象组合成复杂的符合类型的主要方式。在 Go 中，组合是面向对象编程方式的核心。

- JSON
```go
// 把 Go 的数据结构转换成 JSON 称为 marshal。marshal 是通过 json.Marshal 来实现的：
data, err := json.Marshal(movies)
if err != nil {
  log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)

// 为了方便阅读，有一个 json.MarshalIndent 的变体可以输出整齐格式化过的结果
data, err := json.MarshalIndent(movies, "", "  ")
if err != nil {
  log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```

- marshal 使用 Go 结构体成员的名称作为 JSON 对象里面字段的名称（通过反射的方式），只有可导出的成员可以转换成 JSON 字段。
- marshal 的逆操作将 JSON 字符串解码为 Go 数据结构，这个过程叫做 unmarshal，这个是由 json.Unmarshal 实现的。

```go
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
  log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles)
```