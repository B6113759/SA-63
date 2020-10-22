// Code generated by entc, DO NOT EDIT.

package deceasedreceive

const (
	// Label holds the string label denoting the deceasedreceive type in the database.
	Label = "deceased_receive"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeathTime holds the string denoting the death_time field in the database.
	FieldDeathTime = "death_time"

	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "doctor"
	// EdgeRelative holds the string denoting the relative edge name in mutations.
	EdgeRelative = "relative"
	// EdgeCoolroom holds the string denoting the coolroom edge name in mutations.
	EdgeCoolroom = "coolroom"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"

	// Table holds the table name of the deceasedreceive in the database.
	Table = "deceased_receives"
	// DoctorTable is the table the holds the doctor relation/edge.
	DoctorTable = "deceased_receives"
	// DoctorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	DoctorInverseTable = "users"
	// DoctorColumn is the table column denoting the doctor relation/edge.
	DoctorColumn = "user_id"
	// RelativeTable is the table the holds the relative relation/edge.
	RelativeTable = "deceased_receives"
	// RelativeInverseTable is the table name for the Relative entity.
	// It exists in this package in order to avoid circular dependency with the "relative" package.
	RelativeInverseTable = "relatives"
	// RelativeColumn is the table column denoting the relative relation/edge.
	RelativeColumn = "relative_id"
	// CoolroomTable is the table the holds the coolroom relation/edge.
	CoolroomTable = "deceased_receives"
	// CoolroomInverseTable is the table name for the Coolroom entity.
	// It exists in this package in order to avoid circular dependency with the "coolroom" package.
	CoolroomInverseTable = "coolrooms"
	// CoolroomColumn is the table column denoting the coolroom relation/edge.
	CoolroomColumn = "coolroom_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "deceased_receives"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
)

// Columns holds all SQL columns for deceasedreceive fields.
var Columns = []string{
	FieldID,
	FieldDeathTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the DeceasedReceive type.
var ForeignKeys = []string{
	"coolroom_id",
	"patient_id",
	"relative_id",
	"user_id",
}
