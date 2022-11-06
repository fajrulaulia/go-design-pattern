package singleton

import (
	"fmt"
	"log"
	"sync"
)

type Database struct {
	StringConn string
}

func (db *Database) Query() string {
	return fmt.Sprintf("run query = %s", db.StringConn)
}

func (db *Database) Intialize(conn string) {
	db.StringConn = conn
}

var Db *Database
var once sync.Once

func GetDbInstance() *Database {
	if Db == nil {
		log.Println("Create Instance")
		Db = &Database{}
	}
	log.Println("Returning Instance")
	return Db
}

// Singleton Safaly using Once
func GetDbInstanceAsync() *Database {
	once.Do(func() {
		log.Println("Create Instance")
		Db = &Database{}
	})
	log.Println("Returning Instance")
	return Db
}
