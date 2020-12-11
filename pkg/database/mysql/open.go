package mysql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host string
	DatabaseName string
	Port string
	User string
	Password string
}

func NewDatabaseConfig() *Config {
	return &Config{
		Host:         "localhost",
		DatabaseName: "person",
		Port:         "3306",
		User:         "root",
		Password:     "55",
	}
}

func Open(c *Config)	(*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DatabaseName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err =  db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
