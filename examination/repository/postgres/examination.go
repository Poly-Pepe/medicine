package postgres

import (
	"context"
	"errors"
	"time"

	"medicine/domain"

	"github.com/jackc/pgx/v5/pgxpool"
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
	pool *pgxpool.Pool
}

func NewExamination(pool *pgxpool.Pool) domain.ExaminationRepository {
	return &Examination{pool: pool}
}

func (r *Examination) AddDoctor(ctx context.Context, doc *domain.Doctor) error {
	if doc == nil {
		return errDoctorCannotBeNil
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, sqlAddDoctor, doc.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *Examination) AddPatient(ctx context.Context, patient *domain.Patient) error {
	if patient == nil {
		return errPatientCannotBeNil
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(ctx, sqlAddPatient, patient.Name, patient.Gender, patient.BirthDate, patient.Address)
	if err != nil {
		return err
	}

	return nil
}

func (r *Examination) AddMedicine(ctx context.Context, med *domain.Medicine) error {
	if med == nil {
		return errMedicineCannotBeNil
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(ctx, sqlAddMedicine, med.Name, med.MethodOfAdministration, med.Description, med.SideEffects)
	if err != nil {
		return err
	}

	return nil
}

func (r *Examination) AddExamination(ctx context.Context, exam *domain.Examination) error {
	if exam == nil {
		return errExaminationCannotBeNil
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Exec(
		ctx,
		sqlAddExamination,
		exam.PatientID,
		exam.DoctorID,
		exam.ExaminationDate,
		exam.ExaminationPlace,
		exam.Symptoms,
		exam.Diagnosis,
		exam.Prescriptions,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Examination) GetMedicineSideEffects(ctx context.Context, medicineID int) (string, error) {
	if medicineID <= 0 {
		return "", errInvalidMedicineID
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return "", err
	}

	defer conn.Release()

	var sideEffects string
	err = conn.QueryRow(ctx, sqlGetMedicineSideEffects, medicineID).Scan(&sideEffects)
	if err != nil {
		return "", err
	}

	return sideEffects, nil
}

func (r *Examination) GetCountExaminationByDate(ctx context.Context, date time.Time) (int, error) {
	// todo: validate date

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	var count int
	err = conn.QueryRow(ctx, sqlGetCountExaminationByDate, date).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Examination) GetCountExaminationByDiagnosis(ctx context.Context, diagnosis string) (int, error) {
	if diagnosis == "" {
		return 0, errDiagnosisCannotBeEmpty
	}

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	var count int
	err = conn.QueryRow(ctx, sqlGetCountExaminationByDiagnosis, diagnosis).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
