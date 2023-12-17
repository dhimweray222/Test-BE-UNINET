package attendances

import (
	"context"

	"github.com/dhimweray222/test-BE-uninet/model/domain"
)

type AttendanceRepository interface {
	CreateAttendanceTx(ctx context.Context, attendance domain.Attendance) error
	CheckOutTx(ctx context.Context, attendance domain.Attendance) error
	FindOne(ctx context.Context, id string) (domain.Attendance, error)
}
