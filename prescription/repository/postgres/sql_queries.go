package postgres

const (
	sqlAddMedicineToExamination = `
		INSERT INTO prescriptions
		(examination_id, medicine_id)
		VALUES
		($1, $2)`
)
