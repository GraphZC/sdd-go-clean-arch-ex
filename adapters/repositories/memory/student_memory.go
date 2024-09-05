package memory

import (
	"github.com/GraphZC/sdd-go-clean-arch-ex/entities"
	"github.com/GraphZC/sdd-go-clean-arch-ex/repositories"
)

type studentMemory struct {
	students []entities.Student
}

func NewStudentMemory() repositories.StudentRepository {
	return &studentMemory{
		students: []entities.Student{
			{
				ID:   "6410401078",
				Name: "Tanaroeg O-Charoen",
			},
		},
	}
}

func (s *studentMemory) GetAll() []entities.Student {
	return s.students
}
