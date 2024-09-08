package postgres

import (
	"context"

	"medicine/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Prescription struct {
	pool *pgxpool.Pool
}

func NewPrescription(pool *pgxpool.Pool) domain.PrescriptionRepository {
	return &Prescription{pool: pool}
}

func (r *Prescription) AddPrescription(ctx context.Context, examinationID, medicineID int) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(ctx, sqlAddMedicineToExamination, examinationID, medicineID)
	if err != nil {
		return err
	}

	return nil
}
