package person

type Repository interface {
	GetAll() ([]*Person, error)
	GetById(ID uint) (*Person, error)
	Store(person *Person) error
	Update(person *Person) error
	Delete(ID uint) error
}
