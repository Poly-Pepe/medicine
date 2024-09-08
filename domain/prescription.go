package domain

type Prescription struct {
	ID            int
	ExaminationID int
	MedicineID    int
}

type PrescriptionRepository interface {
	AddPrescription(examinationID int, medicineID int) error
}
