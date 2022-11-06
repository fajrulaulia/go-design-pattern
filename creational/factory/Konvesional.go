package factory

import "fmt"

type AccountKonvesional struct {
	Account
}

func (n AccountKonvesional) String() string {
	return fmt.Sprintf("This is a Konvesional Account with Account Number:  %s", n.AccountNo)
}

func createKonvesionalAccount(id int, name string, AccountNo string) AccountIface {
	return &AccountKonvesional{
		Account: Account{
			ID:        id,
			Name:      name,
			AccountNo: AccountNo,
		},
	}
}
