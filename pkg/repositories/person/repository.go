package person

import (
	"database/sql"
	"errors"
	entities "snake/internal/entities/person"
)

type Repository struct {
	DB *sql.DB
}

func NewPersonRepository(DB *sql.DB) *Repository {
	return &Repository{
		DB: DB,
	}
}

func (p Repository) GetAll() ([]*entities.Person, error) {
	var people []*entities.Person

	res, err := p.DB.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var person entities.Person
		err := res.Scan(&person.IDPerson, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		people = append(people, &person)
	}

	return people, nil
}

func (p Repository) GetById(ID uint) (*entities.Person, error) {
	var person entities.Person

	res, err := p.DB.Query("SELECT * FROM person WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(&person.IDPerson, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}
	}

	if person.IDPerson != 0 {
		return &person, nil
	}

	return nil, errors.New("person not found")
}

func (p Repository) Store(person *entities.Person) error {
	stmt, err := p.DB.Prepare("INSERT INTO person (first_name, last_name) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(person.FirstName, person.LastName)
	if err != nil {
		return err
	}

	return nil
}

func (p Repository) Update(person *entities.Person) error {
	stmt, err := p.DB.Prepare("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(person.FirstName, person.LastName, person.IDPerson)
	if err != nil {
		return err
	}

	return nil
}

func (p Repository) Delete(ID uint) error {
	stmt, err := p.DB.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}
