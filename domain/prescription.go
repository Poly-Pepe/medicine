package domain

import "context"

type Prescription struct {
	ID            int
	ExaminationID int
	MedicineID    int
}

type PrescriptionRepository interface {
	AddPrescription(ctx context.Context, examinationID int, medicineID int) error
}
