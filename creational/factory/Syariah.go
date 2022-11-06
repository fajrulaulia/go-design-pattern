package factory

import "fmt"

type AccountSyariah struct {
	Account
}

func (n AccountSyariah) String() string {
	return fmt.Sprintf("This is a Syariah Account with Account Number:  %s", n.AccountNo)
}

func createSyariahAccount(id int, name string, AccountNo string) AccountIface {
	return &AccountKonvesional{
		Account: Account{
			ID:        id,
			Name:      name,
			AccountNo: AccountNo,
		},
	}
}
