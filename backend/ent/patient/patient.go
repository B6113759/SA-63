// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientName holds the string denoting the patient_name field in the database.
	FieldPatientName = "patient_name"
	// FieldPatientAge holds the string denoting the patient_age field in the database.
	FieldPatientAge = "patient_age"

	// EdgeDeceasedreceives holds the string denoting the deceasedreceives edge name in mutations.
	EdgeDeceasedreceives = "deceasedreceives"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// DeceasedreceivesTable is the table the holds the deceasedreceives relation/edge.
	DeceasedreceivesTable = "deceased_receives"
	// DeceasedreceivesInverseTable is the table name for the DeceasedReceive entity.
	// It exists in this package in order to avoid circular dependency with the "deceasedreceive" package.
	DeceasedreceivesInverseTable = "deceased_receives"
	// DeceasedreceivesColumn is the table column denoting the deceasedreceives relation/edge.
	DeceasedreceivesColumn = "patient_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPatientName,
	FieldPatientAge,
}

var (
	// PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	PatientNameValidator func(string) error
	// PatientAgeValidator is a validator for the "patient_age" field. It is called by the builders before save.
	PatientAgeValidator func(int) error
)
