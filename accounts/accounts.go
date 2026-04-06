package accounts

// Account struct
type Account struct {  //lower case -> private, Upper case -> public
	// 싹다 private 로 만들고 생성자 만들기
	owner string
	balance int
}

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

// Balance of receiver's account
// 실제 값을 건들이면 안되는 부분이니까 call by value
func (a Account) Balance() int {
	return a.balance
}