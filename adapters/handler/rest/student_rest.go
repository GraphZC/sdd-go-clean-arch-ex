package rest

import (
	"github.com/GraphZC/sdd-go-clean-arch-ex/usecases"
	"github.com/gofiber/fiber/v2"
)

type StudentRestHandler interface {
	GetAllStudents(c *fiber.Ctx) error
}

type studentRestHandler struct {
	studentUseCase usecases.StudentUseCase
}

func NewStudentRestHandler(studentUseCase usecases.StudentUseCase) StudentRestHandler {
	return &studentRestHandler{
		studentUseCase: studentUseCase,
	}
}

func (s *studentRestHandler) GetAllStudents(c *fiber.Ctx) error {
	students := s.studentUseCase.GetAllStudents()

	return c.Status(fiber.StatusOK).JSON(students)
}
