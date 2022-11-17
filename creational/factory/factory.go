package factory

import "fmt"

func NewAccount(AccountType string, id int, name string, AccountNo string) (AccountIface, error) {
	if AccountType == "syariah" {
		return createSyariahAccount(id, name, AccountNo), nil
	}
	if AccountType == "konvensional" {
		return createKonvesionalAccount(id, name, AccountNo), nil
	}
	return nil, fmt.Errorf("no such account type")
}
