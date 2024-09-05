package repositories

import "github.com/GraphZC/sdd-go-clean-arch-ex/entities"

type StudentRepository interface {
	GetAll() []entities.Student
}
