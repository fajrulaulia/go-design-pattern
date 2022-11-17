package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryBuatAccountKonvesionalSuccess(t *testing.T) {
	acc, err := NewAccount("konvensional", 1, "fajrul", "12345")
	assert.Nil(t, err)
	assert.Equal(t, acc.GetName(), "fajrul")
	assert.Equal(t, acc.GetAccountNo(), "12345")
	assert.Equal(t, acc.GetName(), "fajrul")

	acc.SetAccountNo("54321")
	assert.Equal(t, acc.GetAccountNo(), "54321")

	acc.SetName("Indra")
	assert.Equal(t, acc.GetName(), "Indra")

	acc.SetId(12)
	assert.Equal(t, acc.GetID(), 12)

}

func TestFactoryBuatAccountSyariahSuccess(t *testing.T) {
	acc, err := NewAccount("syariah", 1, "fajrul", "12345")
	assert.Nil(t, err)
	assert.Equal(t, acc.GetName(), "fajrul")
	assert.Equal(t, acc.GetAccountNo(), "12345")
	assert.Equal(t, acc.GetName(), "fajrul")

	acc.SetAccountNo("54321")
	assert.Equal(t, acc.GetAccountNo(), "54321")

	acc.SetName("Indra")
	assert.Equal(t, acc.GetName(), "Indra")

	acc.SetId(12)
	assert.Equal(t, acc.GetID(), 12)

}
