package pointer_err

import (
	"errors"
	"fmt"
)

var insufficetMoneyErr error = errors.New("balance is not enough")

//创建新的计量单位
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//存钱的方法
func (w *Wallet) deposit(money Bitcoin) {
	w.balance += money
}

//查看余额的方法
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	if w.Balance() < money {
		return insufficetMoneyErr
	}
	w.balance -= money
	return nil
}

//实现Bitcoin的stringer方法
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
