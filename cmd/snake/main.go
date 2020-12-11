package main

import (
	"log"
	"snake/pkg/database/mysql"
	"snake/pkg/repositories/person"
	"snake/pkg/webapi"
	"snake/pkg/webapi/setup"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig := mysql.NewDatabaseConfig()
	db, err := mysql.Open(dbConfig)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	personRepository := person.NewPersonRepository(db)

	dependency := setup.Dependency{
		Person: personRepository,
	}

	webapi.ServeAndListen("3456", &dependency)
}
