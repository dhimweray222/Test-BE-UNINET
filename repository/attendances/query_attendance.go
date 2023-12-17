package attendances

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dhimweray222/test-BE-uninet/exception"
	"github.com/dhimweray222/test-BE-uninet/model/domain"
	"github.com/jackc/pgx/v5"
)

func (repository *AttendanceRepositoryImpl) CreateAttendance(ctx context.Context, db pgx.Tx, attendance domain.Attendance) error {
	queryStr := fmt.Sprintf(`SELECT *  FROM %s 
	WHERE user_id = $1 AND DATE_TRUNC('day', clock_in) = CURRENT_DATE`, "attendances")

	data, err := db.Query(context.Background(), queryStr, attendance.UserId)
	log.Println(data)
	if err != nil {
		log.Println(err)
		return err
	}

	defer data.Close()
	log.Println("Hasil Query:", attendance)
	row, err := pgx.CollectOneRow(data, pgx.RowToStructByPos[domain.Attendance])
	if err != nil && err.Error() != "no rows in result set" {
		log.Println("error Find One :", err)
		return err
	}
	log.Println(row)
	check := row.ClockIn.Format("2006-01-01")
	now := time.Now().Format("2006-01-01")
	log.Println(check)
	log.Println(now)
	if now == check {
		return exception.ErrorBadRequest("User Already Checked In")
	}

	log.Print("sini")
	query := fmt.Sprintf(`
	INSERT INTO %s (
		id,
		ip_address,
		latitude,
		longitude,
		clock_in,
		clock_out,
		user_id
	)
	VALUES($1,$2,$3,$4,$5,$6,$7)`, "attendances")

	if _, err := db.Prepare(context.Background(), "create_attendance", query); err != nil {
		return err
	}

	if _, err := db.Exec(context.Background(), "create_attendance",
		attendance.ID,
		attendance.IPAddress,
		attendance.Latitude,
		attendance.Longitude,
		attendance.ClockIn,
		attendance.ClockOut,
		attendance.UserId,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repository *AttendanceRepositoryImpl) FindOneAttendance(ctx context.Context, db pgx.Tx, id string) (domain.Attendance, error) {
	queryStr := fmt.Sprintf(`SELECT *  FROM %s 
	WHERE id = $1`, "attendances")

	attendance, err := db.Query(context.Background(), queryStr, id)

	if err != nil {
		log.Println(err)
		return domain.Attendance{}, err
	}

	defer attendance.Close()
	log.Println("Hasil Query:", attendance)
	data, err := pgx.CollectOneRow(attendance, pgx.RowToStructByPos[domain.Attendance])

	if err != nil {
		log.Println("error Find One :", err)
		return domain.Attendance{}, err
	}

	return data, nil
}
func (repository *AttendanceRepositoryImpl) CheckOut(ctx context.Context, db pgx.Tx, attendance domain.Attendance) error {
	query := fmt.Sprintf("UPDATE %s SET clock_out = $1  WHERE id = $2 ", "attendances")
	_, err := db.Prepare(context.Background(), "clock_out", query)
	if err != nil {
		log.Println("error Find One :", err)
		return err
	}

	data, err := db.Query(context.Background(), "clock_out",
		attendance.ClockOut,
		attendance.ID)

	if err != nil {
		return err
	}

	defer data.Close()

	return nil
}
