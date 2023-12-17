package attendances

import (
	"context"
	"log"

	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/dhimweray222/test-BE-uninet/repository"
	"github.com/jackc/pgx/v5"
)

type AttendanceRepositoryImpl struct {
	DB repository.Store
}

func NewAttendanceRepository(db repository.Store) AttendanceRepository {
	return &AttendanceRepositoryImpl{
		DB: db,
	}
}

func (repository *AttendanceRepositoryImpl) CreateAttendanceTx(ctx context.Context, attendance domain.Attendance) error {
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {
		err = repository.CreateAttendance(ctx, tx, attendance)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
func (repository *AttendanceRepositoryImpl) FindOne(ctx context.Context, id string) (domain.Attendance, error) {
	var err error
	// log.Println("id repo", id)
	var attendance domain.Attendance
	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {
		attendance, err = repository.FindOneAttendance(ctx, tx, id)
		if err != nil {
			// log.Println("sini", err)
			return exception.ErrorBadRequest(err.Error())
		}
		return nil
	})
	if err != nil {
		return domain.Attendance{}, err
	}
	// log.Print("find one")
	return attendance, nil
}

func (repository *AttendanceRepositoryImpl) CheckOutTx(ctx context.Context, attendance domain.Attendance) error {
	var err error

	err = repository.DB.WithTransaction(ctx, func(tx pgx.Tx) error {

		err = repository.CheckOut(ctx, tx, attendance)
		if err != nil {
			log.Println("err checkout")
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Print("checkout")
	return nil
}
