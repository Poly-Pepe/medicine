package postgres

const (
	sqlGetCountExaminationByDate = `
		SELECT COUNT(id) FROM public.examination
		WHERE date = $1`

	sqlGetCountExaminationByDiagnosis = `
		SELECT COUNT(id) FROM public.examination
		WHERE diagnosis = $1`

	sqlGetMedicineSideEffects = `
		SELECT side_effects FROM public.medicines
		WHERE id = $1`

	sqlAddExamination = `
		INSERT INTO public.examination
		(patient_id, doctor_id, examination_date, examination_place, symptoms, diagnosis, prescriptions)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	sqlAddMedicine = `
		INSERT INTO public.medicines
		(name, method_of_administration, description, side_effects)
		values ($1, $2, $3, $4)
		RETURNING id`

	sqlAddDoctor = `
		INSERT INTO public.doctors
		(name)
		VALUES ($1)
		RETURNING id`

	sqlAddPatient = `
		INSERT INTO public.patients
		(name, gender, birth_date, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
)
