package service

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/helper"
	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/dhimweray222/test-BE-uninet/model/web"
	repository "github.com/dhimweray222/test-BE-uninet/repository/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

type UserService interface {
	CreateUser(ctx *fiber.Ctx, request domain.User) (web.UserResponse, error)
	FindAllNotDeletedUser(ctx *fiber.Ctx) ([]web.UserResponse, error)
	FindUserByQuery(ctx *fiber.Ctx, id string, token string) (web.UserResponse, error)
	Login(ctx *fiber.Ctx, request web.LoginRequest) (fiber.Cookie, web.LoginResponse, error)
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) CreateUser(ctx *fiber.Ctx, request domain.User) (web.UserResponse, error) {
	err := service.Validate.Struct(request)

	if err != nil {
		err = helper.ValidateStruct(err)
		return web.UserResponse{}, err
	}

	if !helper.IsEmailValid(request.Email) {
		return web.UserResponse{}, exception.ErrorBadRequest("Not valid email.")
	}

	if len(request.Password) < 6 {
		return web.UserResponse{}, exception.ErrorBadRequest("Password lenght should more then equal 6 character.")
	}

	createdUser := domain.User{
		ID:       request.ID,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,

		CreatedAt: time.Now(),
	}

	createdUser.SetPassword(request.Password)
	createdUser.GenerateID(request.ID)

	userEmail, _ := service.UserRepository.FindUserByQueryTx(ctx.Context(), "email", createdUser.Email)
	if userEmail.ID != "" && userEmail.ID != createdUser.ID {
		return web.UserResponse{}, exception.ErrorBadRequest("Email exist.")
	}

	var createErr error

	err = service.UserRepository.CreateUserTx(ctx.Context(), createdUser)

	if err != nil {
		createErr = err
	}

	if createErr != nil {
		return web.UserResponse{}, createErr
	}
	user, err := service.UserRepository.FindUserByQueryTx(ctx.Context(), "id", createdUser.ID)
	if err != nil || user.ID == "" {
		return web.UserResponse{}, exception.ErrorNotFound("User not exist")
	}

	return domain.ToUserResponse(user), nil
}

func (service *UserServiceImpl) FindUserByQuery(ctx *fiber.Ctx, id string, token string) (web.UserResponse, error) {
	jwtId, jwtEmail, err := helper.ParseJwt(token)
	if err != nil {
		return web.UserResponse{}, err
	}
	log.Println("id : ", jwtId)
	log.Println("email : ", jwtEmail)
	user, err := service.UserRepository.FindUserByQueryTx(ctx.Context(), "id", id)

	if err != nil || user.ID == "" {
		return web.UserResponse{}, exception.ErrorNotFound("User not found.")
	}

	return domain.ToUserResponse(user), nil
}

func (service *UserServiceImpl) FindAllNotDeletedUser(ctx *fiber.Ctx) ([]web.UserResponse, error) {

	users, err := service.UserRepository.FindAllNotDeletedUserTx(ctx.Context())
	if err != nil {
		log.Println("service", err)
	}
	log.Println(users)
	return domain.ToAllUserResponses(users), nil
}

func (service *UserServiceImpl) Login(ctx *fiber.Ctx, request web.LoginRequest) (fiber.Cookie, web.LoginResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return fiber.Cookie{}, web.LoginResponse{}, exception.ErrorBadRequest(err.Error())
	}

	user, err := service.UserRepository.LoginTx(ctx.Context(), request.Email)
	if err != nil || user.ID == "" {
		return fiber.Cookie{}, web.LoginResponse{}, exception.ErrorNotFound("User tidak ditemukan.")
	}

	err = user.ComparePassword(user.Password, request.Password)
	if err != nil {
		return fiber.Cookie{}, web.LoginResponse{}, exception.ErrorBadRequest("Password salah.")
	}

	token, err := helper.GenerateJwt(user.ID, user.Email)
	if err != nil {
		return fiber.Cookie{}, web.LoginResponse{}, exception.ErrorBadRequest(err.Error())
	}

	SessionLogin := os.Getenv("SESSION_LOGIN")

	session, _ := strconv.Atoi(SessionLogin)
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * time.Duration(session)),
		HTTPOnly: true,
	}

	return cookie, domain.ToLoginResponse(user, token), nil
}
