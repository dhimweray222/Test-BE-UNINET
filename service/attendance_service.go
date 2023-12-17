package service

import (
	"time"

	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/helper"
	"github.com/dhimweray222/test-BE-uninet/model/domain"
	repository "github.com/dhimweray222/test-BE-uninet/repository/attendances"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AttendanceServiceImpl struct {
	AttendanceRepository repository.AttendanceRepository
	Validate             *validator.Validate
}

type AttendanceService interface {
	CreateAttendance(ctx *fiber.Ctx, token, ip_address string) (domain.Attendance, error)
	CheckOut(ctx *fiber.Ctx, token, id string) (domain.Attendance, error)
}

func NewAttendanceService(attendanceRepository repository.AttendanceRepository, validate *validator.Validate) AttendanceService {
	return &AttendanceServiceImpl{
		AttendanceRepository: attendanceRepository,
		Validate:             validate,
	}
}

func (service *AttendanceServiceImpl) CreateAttendance(ctx *fiber.Ctx, token, ip_address string) (domain.Attendance, error) {
	jwtId, _, err := helper.ParseJwt(token)
	if err != nil {
		return domain.Attendance{}, err
	}

	// err = service.Validate.Struct(request)
	// if err != nil {
	// 	err = helper.ValidateStruct(err)
	// 	return domain.Attendance{}, err
	// }

	location, err := helper.GetLocation(ip_address)
	if err != nil {
		return domain.Attendance{}, err
	}
	createData := domain.Attendance{
		IPAddress: ip_address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		ClockIn:   time.Now(),
		ClockOut:  time.Time{},
		UserId:    jwtId,
	}

	createData.GenerateID("")

	err = service.AttendanceRepository.CreateAttendanceTx(ctx.Context(), createData)
	if err != nil {
		return domain.Attendance{}, err
	}

	return createData, nil
}

func (service *AttendanceServiceImpl) CheckOut(ctx *fiber.Ctx, token, id string) (domain.Attendance, error) {
	jwtId, _, err := helper.ParseJwt(token)
	var attendance domain.Attendance
	if err != nil {
		return domain.Attendance{}, err
	}

	data := domain.Attendance{
		ID:       id,
		ClockOut: time.Now(),
		UserId:   jwtId,
	}
	attendance, err = service.AttendanceRepository.FindOne(ctx.Context(), id)
	if err != nil {
		return domain.Attendance{}, err
	}
	in := attendance.ClockIn.Format("2006-01-02")
	out := attendance.ClockOut.Format("2006-01-02")
	currentDate := time.Now().Format("2006-01-02")
	if in != currentDate {
		return domain.Attendance{}, exception.ErrorBadRequest("you haven't checked in today")
	}
	if out == currentDate {
		return domain.Attendance{}, exception.ErrorBadRequest("You have checked out today")
	}
	err = service.AttendanceRepository.CheckOutTx(ctx.Context(), data)
	if err != nil {
		return domain.Attendance{}, err
	}
	attendance.ClockOut = data.ClockOut
	return attendance, nil
}
