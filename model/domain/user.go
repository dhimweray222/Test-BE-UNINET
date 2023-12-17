package domain

import (
	"time"

	"github.com/dhimweray222/test-BE-uninet/model/web"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password" `
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) GenerateID(id string) {
	uuid := uuid.New().String()
	user.ID = uuid
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hashedPassword)
}

func (user *User) ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ToUserResponse(user User) web.UserResponse {
	return web.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
func ToAllUserResponses(users []User) []web.UserResponse {
	var datausers []web.UserResponse
	for _, data := range users {
		datausers = append(datausers, ToUserResponse(data))
	}
	return datausers
}

func ToLoginResponse(user User, token string) web.LoginResponse {
	return web.LoginResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
