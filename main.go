package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhimweray222/test-BE-uninet/config"
	"github.com/dhimweray222/test-BE-uninet/controller"
	"github.com/dhimweray222/test-BE-uninet/repository"
	attendanceRepo "github.com/dhimweray222/test-BE-uninet/repository/attendances"
	userRepo "github.com/dhimweray222/test-BE-uninet/repository/users"
	"github.com/dhimweray222/test-BE-uninet/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	time.Local = time.UTC
	validate := validator.New()
	db := config.NewPostgresPool()
	userStore := repository.NewUserStore(db)

	// user
	userRepository := userRepo.NewUserRepository(userStore)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	// attendances
	attendanceRepository := attendanceRepo.NewAttendanceRepository(userStore)
	attendanceService := service.NewAttendanceService(attendanceRepository, validate)
	attendanceController := controller.NewAttendanceController(attendanceService)

	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	//define routes
	userController.NewUserRouter(app)
	attendanceController.NewAttendanceRouter(app)

	host := fmt.Sprintf("%s:%s", os.Getenv("SERVER_URI"), os.Getenv("SERVER_PORT"))
	err := app.Listen(host)

	log.Println(err)
}
