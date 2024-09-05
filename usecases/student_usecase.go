package usecases

import (
	"github.com/GraphZC/sdd-go-clean-arch-ex/entities"
	"github.com/GraphZC/sdd-go-clean-arch-ex/repositories"
)

type StudentUseCase interface {
	GetAllStudents() []entities.Student
}

type studentService struct {
	studentRepo repositories.StudentRepository
}

func NewStudentService(studentRepo repositories.StudentRepository) StudentUseCase {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (s *studentService) GetAllStudents() []entities.Student {
	return s.studentRepo.GetAll()
}
