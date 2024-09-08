package front

import (
	"context"
	"path/filepath"
	"strconv"
	"time"

	"medicine/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	log "github.com/sirupsen/logrus"
)

type Examination struct {
	ExaminationUseCase domain.ExaminationUseCase
	App                *fiber.App
}

func NewExamination(uc domain.ExaminationUseCase) *Examination {
	path, err := filepath.Abs("./examination/delivery/front/views")
	if err != nil {
		log.Fatal("Не удалось определить путь к шаблонам")
	}

	examination := &Examination{
		ExaminationUseCase: uc,
		App: fiber.New(
			fiber.Config{
				Views: html.New(path, ".html"),
			},
		),
	}

	examination.App.Get("/", examination.GetIndex)

	examination.App.Get("/add-doctor", examination.GetAddDoctorPage)
	examination.App.Get("/add-patient", examination.GetAddPatientPage)
	examination.App.Get("/add-medicine", examination.GetAddMedicinePage)
	examination.App.Get("/add-examination", examination.GetAddExaminationPage)

	examination.App.Post("/add-doctor", examination.PostAddDoctor)
	examination.App.Post("/add-patient", examination.PostAddPatient)
	examination.App.Post("/add-medicine", examination.PostAddMedicine)
	examination.App.Post("/add-examination", examination.PostAddExamination)

	examination.App.Get("/medicine-side-effects/:id", examination.GetMedicineSideEffects)
	examination.App.Get("/count-examinations-by-date/:date", examination.GetCountExaminationsByDate)
	examination.App.Get("/count-examinations-by-diagnosis/:diagnosis", examination.GetCountExaminationsByDiagnosis)

	return examination
}

func (e *Examination) GetIndex(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func (e *Examination) GetAddDoctorPage(c *fiber.Ctx) error {
	return c.Render("add_doctor", nil)
}

func (e *Examination) GetAddPatientPage(c *fiber.Ctx) error {
	return c.Render("add_patient", nil)
}

func (e *Examination) GetAddMedicinePage(c *fiber.Ctx) error {
	return c.Render("add_medicine", nil)
}

func (e *Examination) GetAddExaminationPage(c *fiber.Ctx) error {
	return c.Render("add_examination", nil)
}

func (e *Examination) PostAddDoctor(c *fiber.Ctx) error {
	doctor := domain.Doctor{
		Name: c.FormValue("name"),
	}

	err := e.ExaminationUseCase.AddDoctor(context.Background(), &doctor)
	if err != nil {
		return c.Status(500).SendString("Ошибка при добавлении врача")
	}

	return c.SendString("Доктор успешно добавлен!")
}

func (e *Examination) PostAddPatient(c *fiber.Ctx) error {
	patient := domain.Patient{
		Name:      c.FormValue("name"),
		Gender:    c.FormValue("gender"),
		Address:   c.FormValue("address"),
		BirthDate: time.Now(),
	}

	err := e.ExaminationUseCase.AddPatient(context.Background(), &patient)
	if err != nil {
		return c.Status(500).SendString("Ошибка при добавлении пациента")
	}

	return c.SendString("Пациент успешно добавлен!")
}

func (e *Examination) PostAddMedicine(c *fiber.Ctx) error {
	medicine := domain.Medicine{
		Name:                   c.FormValue("name"),
		MethodOfAdministration: c.FormValue("method"),
		Description:            c.FormValue("description"),
		SideEffects:            c.FormValue("side_effects"),
	}

	err := e.ExaminationUseCase.AddMedicine(context.Background(), &medicine)
	if err != nil {
		return c.Status(500).SendString("Ошибка при добавлении лекарства")
	}

	return c.SendString("Лекарство успешно добавлено!")
}

func (e *Examination) PostAddExamination(c *fiber.Ctx) error {
	docID, err := strconv.Atoi(c.FormValue("doctor_id"))
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении ID врача")
	}

	patID, err := strconv.Atoi(c.FormValue("patient_id"))
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении ID пациента")
	}

	exam := domain.Examination{
		DoctorID:         docID,
		PatientID:        patID,
		ExaminationDate:  time.Now(),
		ExaminationPlace: c.FormValue("examination_place"),
		Symptoms:         c.FormValue("symptoms"),
		Diagnosis:        c.FormValue("diagnosis"),
		Prescriptions:    c.FormValue("prescriptions"),
	}

	err = e.ExaminationUseCase.AddExamination(context.Background(), &exam)
	if err != nil {
		return c.Status(500).SendString("Ошибка при добавлении осмотра")
	}

	return c.SendString("Осмотр успешно добавлен!")
}

func (e *Examination) GetMedicineSideEffects(c *fiber.Ctx) error {
	medicineID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении ID лекарства")
	}

	sideEffects, err := e.ExaminationUseCase.GetMedicineSideEffects(context.Background(), medicineID)
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении побочных эффектов")
	}

	return c.Render("medicine_side_effects", sideEffects)
}

func (e *Examination) GetCountExaminationsByDate(c *fiber.Ctx) error {
	date, err := time.Parse("2006-01-02", c.Params("date"))
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении даты")
	}

	count, err := e.ExaminationUseCase.GetCountExaminationByDate(context.Background(), date)
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении данных")
	}

	return c.Render("count_examinations_by_date", fiber.Map{
		"Date":  date,
		"Count": count,
	})
}

func (e *Examination) GetCountExaminationsByDiagnosis(c *fiber.Ctx) error {
	diagnosis := c.Params("diagnosis")

	count, err := e.ExaminationUseCase.GetCountExaminationByDiagnosis(context.Background(), diagnosis)
	if err != nil {
		return c.Status(500).SendString("Ошибка при получении данных")
	}

	return c.Render("count_examinations_by_diagnosis", fiber.Map{
		"Diagnosis": diagnosis,
		"Count":     count,
	})
}
