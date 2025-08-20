package dbase

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

type DB struct {
	TypeDB   string
	Host     string
	Name     string
	User     string
	Password string
	SSL      string
}

func NewDb(typeDB, host, name, user, password string) *DB {
	return &DB{
		TypeDB: typeDB,
		Host:   host,
        User: user,
        Password: password,
	}
}

func (db *DB) GetConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", db.TypeDB, db.User, db.Password, db.Host, db.Name)
}
