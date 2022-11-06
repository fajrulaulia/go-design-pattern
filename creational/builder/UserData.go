package builder

import "time"

type UserData struct {
	ID         int
	Name       string
	Birthdate  time.Time
	BirthPlace string
	MotherName string
	Address    string
	Branch     string
}
