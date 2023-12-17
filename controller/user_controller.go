package controller

import (
	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/dhimweray222/test-BE-uninet/model/web"
	"github.com/dhimweray222/test-BE-uninet/service"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

type UserControllerImpl struct {
	UserService service.UserService
}

type UserController interface {
	NewUserRouter(app *fiber.App)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) NewUserRouter(app *fiber.App) {
	user := app.Group("/users")
	user.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	user.Post("/", controller.CreateUser)
	user.Get("/:id", controller.GetUserById)
	user.Get("/", controller.GetAllUsers)
	user.Post("/login", controller.Login)

}

func (controller *UserControllerImpl) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	cookieValue := ctx.Cookies("token")

	response, err := controller.UserService.FindUserByQuery(ctx, id, cookieValue)

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

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	var request web.LoginRequest
	_ = ctx.BodyParser(&request)

	cookie, resp, err := controller.UserService.Login(ctx, request)

	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	ctx.Cookie(&cookie)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    resp,
	})
}

func (controller *UserControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	var request domain.User

	if err := ctx.BodyParser(&request); err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	response, err := controller.UserService.CreateUser(ctx, request)
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

func (controller *UserControllerImpl) GetAllUsers(ctx *fiber.Ctx) error {

	response, err := controller.UserService.FindAllNotDeletedUser(ctx)
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
