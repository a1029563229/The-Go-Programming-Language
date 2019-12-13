# Concurrency with Shared Variables

## Race Conditions

- 在普通的单线程项目中，我们的函数都是按顺序执行的；在多线程的项目中，我们无法确保函数的执行顺序；
- 一个包级别的导出函数应该是并发安全类型的函数，因为一个并发的函数有太多的原因导致它不能正常工作，比如死锁和资源耗尽。我们不可能去逐一处理这些问题，所以我们应该把重点放在 Race Condition 上。

## Lock/Unlock

- 死锁问题

```go
var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// 重复调用 mu.Lock() 将导致死锁问题
func Withdraw(amount int) bool {
  mu.Lock()
  defer mu.Unlock()
  Deposit(-amount)
  if Balance() < 0 {
    Deposit(amount)
    return false // insufficient funds
  }
  return true
}
```

- 解决死锁问题

```go
// 将问题进行进一步的原子级拆分
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    deposit(-amount)
    if balance < 0 {
        deposit(amount)
        return false // insufficient funds
    }
    return true
}

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    deposit(amount)
}

func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }
```

## 读写锁

- 读的操作
```go
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}
```