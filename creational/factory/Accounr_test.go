package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryBuatAccountKonvesionalSuccess(t *testing.T) {
	acc, err := newAccount("syariah", 1, "fajrul", "12345")

	assert.Nil(t, err)

	assert.Equal(t, acc.GetName(), "fajrul")
	assert.Equal(t, acc.GetAccountNo(), "12345")
	assert.Equal(t, acc.GetName(), "fajrul")

}
