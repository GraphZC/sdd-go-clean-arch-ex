package main

import (
	"github.com/GraphZC/sdd-go-clean-arch-ex/adapters/handler/rest"
	"github.com/GraphZC/sdd-go-clean-arch-ex/adapters/repositories/memory"
	"github.com/GraphZC/sdd-go-clean-arch-ex/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	studentRepo := memory.NewStudentMemory()
	studentService := usecases.NewStudentService(studentRepo)
	studentHandler := rest.NewStudentRestHandler(studentService)

	app.Get("/students", studentHandler.GetAllStudents)

	app.Listen(":8000")
}
