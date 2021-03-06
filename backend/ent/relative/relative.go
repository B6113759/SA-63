// Code generated by entc, DO NOT EDIT.

package relative

const (
	// Label holds the string label denoting the relative type in the database.
	Label = "relative"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRelativeName holds the string denoting the relative_name field in the database.
	FieldRelativeName = "relative_name"

	// EdgeDeceasedreceives holds the string denoting the deceasedreceives edge name in mutations.
	EdgeDeceasedreceives = "deceasedreceives"

	// Table holds the table name of the relative in the database.
	Table = "relatives"
	// DeceasedreceivesTable is the table the holds the deceasedreceives relation/edge.
	DeceasedreceivesTable = "deceased_receives"
	// DeceasedreceivesInverseTable is the table name for the DeceasedReceive entity.
	// It exists in this package in order to avoid circular dependency with the "deceasedreceive" package.
	DeceasedreceivesInverseTable = "deceased_receives"
	// DeceasedreceivesColumn is the table column denoting the deceasedreceives relation/edge.
	DeceasedreceivesColumn = "relative_id"
)

// Columns holds all SQL columns for relative fields.
var Columns = []string{
	FieldID,
	FieldRelativeName,
}

var (
	// RelativeNameValidator is a validator for the "relative_name" field. It is called by the builders before save.
	RelativeNameValidator func(string) error
)
