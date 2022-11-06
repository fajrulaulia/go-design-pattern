package builder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilderdUserDataCreateNasabahBaruSuccess(t *testing.T) {

	var nasabahBaru = NewUserDataBuilder()
	nasabahBaru.SetID(1)
	nasabahBaru.SetName("Fajrul")
	nasabahBaru.SetBirthdate("1999-11-25")
	nasabahBaru.SetBirthPlace("Banda Aceh")
	nasabahBaru.SetMotherName("Rahmah binti Luqman")
	nasabahBaru.SetAddress("Jln Teuku Imum Lueng Bata")
	nasabahBaru.SetBranch("MDN-1")
	result, err := nasabahBaru.Build()

	assert.Nil(t, err)
	assert.Equal(t, result.Name, "Fajrul")
	assert.Equal(t, result.MotherName, "Rahmah binti Luqman")
	assert.Equal(t, result.Birthdate.Day(), 25)
	assert.Equal(t, result.Birthdate.Month(), time.Month(11))
	assert.Equal(t, result.Birthdate.Year(), 1999)

}

func TestBuilderdUserDataCreateNasabahBaruBranchNotFound(t *testing.T) {

	var nasabahBaru = NewUserDataBuilder()
	nasabahBaru.SetID(1)
	nasabahBaru.SetName("Fajrul")
	nasabahBaru.SetBirthdate("1999-11-25")
	nasabahBaru.SetBirthPlace("Banda Aceh")
	nasabahBaru.SetMotherName("Rahmah binti Luqman")
	nasabahBaru.SetAddress("Jln Teuku Imum Lueng Bata")
	nasabahBaru.SetBranch("KKKK-1")
	result, err := nasabahBaru.Build()

	assert.NotNil(t, err)
	assert.Nil(t, result)

}
