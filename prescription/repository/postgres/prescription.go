package postgres

import (
	"medicine/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Prescription struct {
	pool *pgxpool.Pool
}

func NewPrescription(pool *pgxpool.Pool) domain.PrescriptionRepository {
	return &Prescription{pool: pool}
}

func (r *Prescription) AddPrescription(examinationID int, medicineID int) error {
	// todo: fix
	return nil
}
