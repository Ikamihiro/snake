package setup

import "snake/pkg/repositories/person"

type Dependency struct {
	Person *person.Repository
}
