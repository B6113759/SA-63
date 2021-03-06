// Code generated by entc, DO NOT EDIT.

package coolroom

const (
	// Label holds the string label denoting the coolroom type in the database.
	Label = "coolroom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCoolroomName holds the string denoting the coolroom_name field in the database.
	FieldCoolroomName = "coolroom_name"

	// EdgeDeceasedreceives holds the string denoting the deceasedreceives edge name in mutations.
	EdgeDeceasedreceives = "deceasedreceives"
	// EdgeCoolroomtype holds the string denoting the coolroomtype edge name in mutations.
	EdgeCoolroomtype = "coolroomtype"

	// Table holds the table name of the coolroom in the database.
	Table = "coolrooms"
	// DeceasedreceivesTable is the table the holds the deceasedreceives relation/edge.
	DeceasedreceivesTable = "deceased_receives"
	// DeceasedreceivesInverseTable is the table name for the DeceasedReceive entity.
	// It exists in this package in order to avoid circular dependency with the "deceasedreceive" package.
	DeceasedreceivesInverseTable = "deceased_receives"
	// DeceasedreceivesColumn is the table column denoting the deceasedreceives relation/edge.
	DeceasedreceivesColumn = "coolroom_id"
	// CoolroomtypeTable is the table the holds the coolroomtype relation/edge.
	CoolroomtypeTable = "coolrooms"
	// CoolroomtypeInverseTable is the table name for the CoolroomType entity.
	// It exists in this package in order to avoid circular dependency with the "coolroomtype" package.
	CoolroomtypeInverseTable = "coolroom_types"
	// CoolroomtypeColumn is the table column denoting the coolroomtype relation/edge.
	CoolroomtypeColumn = "coolroomtype_id"
)

// Columns holds all SQL columns for coolroom fields.
var Columns = []string{
	FieldID,
	FieldCoolroomName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Coolroom type.
var ForeignKeys = []string{
	"coolroomtype_id",
}

var (
	// CoolroomNameValidator is a validator for the "coolroom_name" field. It is called by the builders before save.
	CoolroomNameValidator func(string) error
)
