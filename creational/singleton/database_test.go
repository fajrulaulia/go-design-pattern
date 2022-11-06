package singleton

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingletonInitialzieDatabaseSuccess(t *testing.T) {
	initDb := GetDbInstance()
	assert.Equal(t, initDb, &Database{StringConn: ""})
	initDb.Intialize("10.10.10.10:1000-/db")
	assert.Equal(t, initDb, &Database{StringConn: "10.10.10.10:1000-/db"})
	result := initDb.Query()
	assert.Equal(t, result, "run query = 10.10.10.10:1000-/db")

	initDb2 := GetDbInstance()
	assert.Equal(t, initDb, &Database{StringConn: "10.10.10.10:1000-/db"})
	assert.Equal(t, initDb.Query(), "run query = 10.10.10.10:1000-/db")

	initDb2.Intialize("10.10.10.20:1000-/db")
	assert.Equal(t, initDb2, &Database{StringConn: "10.10.10.20:1000-/db"})
	result2 := initDb2.Query()
	assert.Equal(t, result2, "run query = 10.10.10.20:1000-/db")

}

func TestSingletonInitialzieDatabaseAsyncSuccess(t *testing.T) {

	for i := 0; i < 100; i++ {
		go GetDbInstanceAsync()
	}
	time.Sleep(time.Second * 10)
}
