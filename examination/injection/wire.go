//go:build wire
// +build wire

package injection

import (
	"medicine/examination/delivery/front"
	examinationRepo "medicine/examination/repository/postgres"
	examinationUseCase "medicine/examination/usecase"
	prescriptionRepo "medicine/prescription/repository/postgres"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeExaminationFront(
	pool *pgxpool.Pool,
) *front.Examination {
	wire.Build(
		examinationRepo.NewExamination,
		prescriptionRepo.NewPrescription,
		examinationUseCase.NewExamination,
		front.NewExamination,
	)

	return &front.Examination{}
}
