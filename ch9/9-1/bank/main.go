// Package bank 提供了并发安全的银行账户
package bank

var deposits = make(chan int) // 发送存款金额的通道
var balances = make(chan int) // 接收账户余额的通道

// Deposit 存款函数 (通过通道发送存款请求)
func Deposit(amount int) { deposits <- amount }

// Balance 查询余额函数 (通过通道接收余额)
func Balance() int { return <-balances }

func teller() { // 银行柜员 goroutine (专门处理账户操作)
    var balance int // 账户余额 (只在 teller goroutine 内部使用，不共享！)

    for { // 循环处理请求
        select { // 监听两个通道
        case amount := <-deposits:  // 收到存款请求
            balance += amount
        case balances <- balance:   // 收到查询余额请求
            // 发送当前余额到 balances 通道
        }
    }
}

func init() { // 初始化函数，启动 teller goroutine
    go teller() // 启动银行柜员 goroutine
}
