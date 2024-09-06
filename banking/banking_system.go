package banking_system

import (
	"errors"
	"fmt"
)

func main() {
	account := NewAccount("Roy")
	account.Deposit(10)
	fmt.Println(account) //account만 출력하여 String()를 호출
}

// Account 객체
type Account struct {
	owner   string
	balance int
}

// 계좌 생성(balance는 0으로 고정, owner만 입력받는다.)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 계좌 잔액 반환
func (a Account) Balance() int {
	return a.balance
}

// 계좌 주인 반환
func (a Account) Owner() string {
	return a.owner
}

// account를 호출하면 자동으로 String() Method를 호출함. 파이썬의 __str__같은 개념.
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account has ", a.Balance())
}

// 입금
func (a *Account) Deposit(amount int) { //account의 메소드
	a.balance += amount
}

// 출금
func (a Account) Withdraw(amount int) error { //Go에서 error는 error와 nil 두 가지로 반환이 가능
	if a.balance < amount {
		return errNomoney //잔액 부족 에러 호출
	}
	a.balance -= amount
	return nil //nil=null
}

var errNomoney = errors.New("Can't withdraw") //출금 시 잔액 부족 대응 에러 생성

// 계좌 주인 변경
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner //a.owner(기존 owner)를 newOwner로 교체
}
