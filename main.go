package main

import (
	"log"

	c "github.com/fajrulaulia/golang-design-pattern/creational/factory"
)

func main() {

	accountSyariah, _ := c.NewAccount("syariah", 1, "fajrul aulia", "111111")

	pubDetails(accountSyariah)
}

func pubDetails(acc c.AccountIface) {
	log.Printf("--------------------\n")
	log.Printf("%s\n", acc)
	log.Printf("Type: %T\n", acc)
	log.Printf("Name: %d\n", acc.GetID())
	log.Printf("Pages: %s\n", acc.GetAccountNo())
	log.Printf("Name: %s\n", acc.GetName())
	log.Printf("--------------------\n")
}
