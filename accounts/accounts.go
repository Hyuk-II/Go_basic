package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {  //lower case -> private, Upper case -> public
	// 싹다 private 로 만들고 생성자 만들기
	owner string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}

	return &account // 주소 반환
}

// method, func receiver funcName arg ...
// receiver -> 객체
// Deposit x amount on receiver's account
// 실제 값을 수정해야 하니까 call by reference
func (a *Account) Deposit(amount int) { 
	a.balance += amount
}

// Withdra amount
// try-catch 없어서 매번 리턴 확인
func (a *Account) WithDraw(amount int) error {
	
	if a.balance < amount  {
		return errNoMoney
	}
	a.balance -= amount
	return nil // null error, 에러 없다는 뜻
}


// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Balance of receiver's account
// 실제 값을 건들이면 안되는 부분이니까 call by value
func (a Account) Balance() int {
	return a.balance
}

// Owner of receiver
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(),"'s Account\nHas: ",a.Balance())
}