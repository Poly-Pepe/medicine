package domain

import (
	"context"
	"time"
)

type Examination struct {
	ID               int
	PatientID        int
	DoctorID         int
	ExaminationDate  time.Time
	ExaminationPlace string
	Symptoms         string
	Diagnosis        string
	Prescriptions    string
}

type Doctor struct {
	ID   int
	Name string
}

type Patient struct {
	ID        int
	Name      string
	Gender    string
	BirthDate time.Time
	Address   string
}

type Medicine struct {
	ID                     int
	Name                   string
	MethodOfAdministration string
	Description            string
	SideEffects            string
}

type ExaminationUseCase interface {
	AddDoctor(ctx context.Context, doc *Doctor) error
	AddPatient(ctx context.Context, patient *Patient) error
	AddMedicine(ctx context.Context, med *Medicine) error
	AddExamination(ctx context.Context, exam *Examination) error

	ListDoctors(ctx context.Context) ([]*Doctor, error)
	ListPatients(ctx context.Context) ([]*Patient, error)
	ListMedicines(ctx context.Context) ([]*Medicine, error)

	GetMedicineSideEffects(ctx context.Context, medicineID int) (string, error)
	GetCountExaminationByDate(ctx context.Context, date time.Time) (int, error)
	GetCountExaminationByDiagnosis(ctx context.Context, diagnosis string) (int, error)

	AddPrescription(ctx context.Context, examinationID int, medicineID int) error
}

type ExaminationRepository interface {
	AddDoctor(ctx context.Context, doc *Doctor) error
	AddPatient(ctx context.Context, patient *Patient) error
	AddMedicine(ctx context.Context, med *Medicine) error
	AddExamination(ctx context.Context, exam *Examination) error

	ListDoctors(ctx context.Context) ([]*Doctor, error)
	ListPatients(ctx context.Context) ([]*Patient, error)
	ListMedicines(ctx context.Context) ([]*Medicine, error)

	GetMedicineSideEffects(ctx context.Context, medicineID int) (string, error)
	GetCountExaminationByDate(ctx context.Context, date time.Time) (int, error)
	GetCountExaminationByDiagnosis(ctx context.Context, diagnosis string) (int, error)
}
