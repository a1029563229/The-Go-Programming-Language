# 反射

Go 语言提供了一种机制，在编译时不知道类型的情况下，可更新变量、在运行时查看值、调用方法以及直接对它们的布局进行操作，这种机制称为反射（reflection）。

反射功能由 `reflect` 包提供，它定义了两个重要的类型：`Type` 和 `Value`。`Type` 表示 Go 语言的一个类型，它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分，比如一个结构的各个字段或者一个函数的各个参数。`reflect.Type` 接口只有一个实现，即类型描述符，接口值中的动态类型也是类型描述符。

`reflect.Type` 函数接受任何的 `interface{}` 参数，并且把接口中的动态类型以 `reflect.Type` 形式返回：

```go
t := reflect.Typeof(3) // 一个 reflect.Type
fmt.Println(t.String()) // "int"
fmt.Println(t)          // "int"
```

把一个具体指赋给一个接口类型时会发送一个隐式类型转换，转换会生成一个包含两部分内容的接口值：动态类型部分是操作数的类型（int），动态值部分是操作数的值（3）。

使用反射将任何值格式化为一个字符串：

```go
// Any 将任何值格式化为一个字符串
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom 格式化一个值，且不分析它的内部结构
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ... 为简化起见，省略了浮点数和复数的分支...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // 1
	fmt.Println(Any(d))                  // 1
	fmt.Println(Any([]int64{x}))         // []int64 0xc00009a020
	fmt.Println(Any([]time.Duration{d})) // []time.Duration 0xc00009a028
}
```