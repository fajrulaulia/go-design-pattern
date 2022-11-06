package factory

type AccountIface interface {
	SetId(ID int)
	SetName(name string)
	SetAccountNo(accountNo string)
	GetID() int
	GetName() string
	GetAccountNo() string
}

type Account struct {
	ID        int
	Name      string
	AccountNo string
}

func (a *Account) SetId(ID int) {
	a.ID = ID
}

func (a *Account) SetName(name string) {
	a.Name = name
}
func (a *Account) SetAccountNo(accountNo string) {
	a.AccountNo = accountNo
}

func (a *Account) GetID() int {
	return a.ID
}
func (a *Account) GetName() string {
	return a.Name
}
func (a *Account) GetAccountNo() string {
	return a.AccountNo
}
