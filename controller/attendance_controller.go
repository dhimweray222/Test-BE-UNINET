package controller

import (
	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/model/web"
	"github.com/dhimweray222/test-BE-uninet/service"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

type AttendanceControllerImpl struct {
	AttendanceService service.AttendanceService
}

type AttendanceController interface {
	NewAttendanceRouter(app *fiber.App)
}

func NewAttendanceController(attendanceService service.AttendanceService) AttendanceController {
	return &AttendanceControllerImpl{
		AttendanceService: attendanceService,
	}
}

func (controller *AttendanceControllerImpl) NewAttendanceRouter(app *fiber.App) {
	attendance := app.Group("/attendances")
	attendance.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	attendance.Post("/", controller.CreateAttendance)
	attendance.Put("/:id", controller.CheckOut)

}

func (controller *AttendanceControllerImpl) CreateAttendance(ctx *fiber.Ctx) error {

	// clientIP := ctx.IP()
	clientIP := "36.78.32.1"
	cookieValue := ctx.Cookies("token")

	if cookieValue == "" {
		return exception.ErrorHandler(ctx, exception.ErrorBadRequest("You haven't logged in"))
	}
	response, err := controller.AttendanceService.CreateAttendance(ctx, cookieValue, clientIP)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    response,
	})
}

func (controller *AttendanceControllerImpl) CheckOut(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	cookieValue := ctx.Cookies("token")
	if cookieValue == "" {
		return exception.ErrorHandler(ctx, exception.ErrorBadRequest("You haven't logged in"))
	}
	response, err := controller.AttendanceService.CheckOut(ctx, cookieValue, id)
	if err != nil {

		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    response,
	})
}
