package usecase

import (
	"context"
	"errors"
	"time"

	"medicine/domain"
)

var (
	errDoctorCannotBeNil      = errors.New("doctor cannot be nil")
	errPatientCannotBeNil     = errors.New("patient cannot be nil")
	errMedicineCannotBeNil    = errors.New("medicine cannot be nil")
	errExaminationCannotBeNil = errors.New("examination cannot be nil")
	errInvalidMedicineID      = errors.New("invalid medicine ID")
	errDiagnosisCannotBeEmpty = errors.New("diagnosis cannot be empty")
)

type Examination struct {
	ExaminationRepo  domain.ExaminationRepository
	PrescriptionRepo domain.PrescriptionRepository
}

func NewExamination(
	examinationRepo domain.ExaminationRepository,
	prescriptionRepo domain.PrescriptionRepository,
) domain.ExaminationUseCase {
	return &Examination{
		ExaminationRepo:  examinationRepo,
		PrescriptionRepo: prescriptionRepo,
	}
}

func (e *Examination) AddDoctor(ctx context.Context, doc *domain.Doctor) error {
	if doc == nil {
		return errDoctorCannotBeNil
	}

	return e.ExaminationRepo.AddDoctor(ctx, doc)
}

func (e *Examination) AddPatient(ctx context.Context, patient *domain.Patient) error {
	if patient == nil {
		return errPatientCannotBeNil
	}

	return e.ExaminationRepo.AddPatient(ctx, patient)
}

func (e *Examination) AddMedicine(ctx context.Context, med *domain.Medicine) error {
	if med == nil {
		return errMedicineCannotBeNil
	}

	return e.ExaminationRepo.AddMedicine(ctx, med)
}

func (e *Examination) AddExamination(ctx context.Context, exam *domain.Examination) error {
	if exam == nil {
		return errExaminationCannotBeNil
	}

	return e.ExaminationRepo.AddExamination(ctx, exam)
}

func (e *Examination) GetMedicineSideEffects(ctx context.Context, medicineID int) (string, error) {
	if medicineID <= 0 {
		return "", errInvalidMedicineID
	}

	return e.ExaminationRepo.GetMedicineSideEffects(ctx, medicineID)
}

func (e *Examination) GetCountExaminationByDate(ctx context.Context, date time.Time) (int, error) {
	// todo: validate date
	return e.ExaminationRepo.GetCountExaminationByDate(ctx, date)
}

func (e *Examination) GetCountExaminationByDiagnosis(ctx context.Context, diagnosis string) (int, error) {
	if diagnosis == "" {
		return 0, errDiagnosisCannotBeEmpty
	}

	return e.ExaminationRepo.GetCountExaminationByDiagnosis(ctx, diagnosis)
}

func (e *Examination) ListDoctors(ctx context.Context) ([]*domain.Doctor, error) {
	return e.ExaminationRepo.ListDoctors(ctx)
}

func (e *Examination) ListPatients(ctx context.Context) ([]*domain.Patient, error) {
	return e.ExaminationRepo.ListPatients(ctx)
}

func (e *Examination) ListMedicines(ctx context.Context) ([]*domain.Medicine, error) {
	return e.ExaminationRepo.ListMedicines(ctx)
}

func (e *Examination) AddPrescription(ctx context.Context, examinationID int, medicineID int) error {
	return e.PrescriptionRepo.AddPrescription(ctx, examinationID, medicineID)
}
