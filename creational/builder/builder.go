package builder

import (
	"fmt"
	"time"
)

type UserDataBuilder struct {
	ID         int
	Name       string
	Birthdate  time.Time
	BirthPlace string
	MotherName string
	Address    string
	Branch     string
}

func NewUserDataBuilder() *UserDataBuilder {
	return &UserDataBuilder{}
}

func (userDataBuilder *UserDataBuilder) SetID(ID int) {
	userDataBuilder.ID = ID
}

func (userDataBuilder *UserDataBuilder) SetName(Name string) {
	userDataBuilder.Name = Name
}

func (userDataBuilder *UserDataBuilder) SetBirthdate(birthdate string) {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, birthdate+"T00:00:00.000Z")
	if err != nil {
		panic(err)
	}
	userDataBuilder.Birthdate = t
}

func (userDataBuilder *UserDataBuilder) SetBirthPlace(birthplace string) {
	userDataBuilder.BirthPlace = birthplace
}

func (userDataBuilder *UserDataBuilder) SetMotherName(motherName string) {
	userDataBuilder.MotherName = motherName
}

func (userDataBuilder *UserDataBuilder) SetAddress(address string) {
	userDataBuilder.Address = address
}

func (userDataBuilder *UserDataBuilder) SetBranch(branch string) {
	userDataBuilder.Branch = branch
}

func (userDataBuilder *UserDataBuilder) Build() (*UserData, error) {

	branch := ""
	switch userDataBuilder.Branch {
	case "WB-1":
		branch = "WB-1 - Warung Buncit, jakarta Selatan"

	case "MDN-1":
		branch = "MDN-1 - Medan Balaikota, Medan"

	}
	if branch == "" {
		return nil, fmt.Errorf("error: branch is empty/not registered")

	}

	return &UserData{
		ID:         userDataBuilder.ID,
		Name:       userDataBuilder.Name,
		Birthdate:  userDataBuilder.Birthdate,
		BirthPlace: userDataBuilder.BirthPlace,
		MotherName: userDataBuilder.MotherName,
		Address:    branch,
		Branch:     userDataBuilder.Branch,
	}, nil
}
