// Code generated by entc, DO NOT EDIT.

package stadium

const (
	// Label holds the string label denoting the stadium type in the database.
	Label = "stadium"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNamestadium holds the string denoting the namestadium field in the database.
	FieldNamestadium = "namestadium"

	// EdgeStadiumConfirmation holds the string denoting the stadiumconfirmation edge name in mutations.
	EdgeStadiumConfirmation = "StadiumConfirmation"

	// Table holds the table name of the stadium in the database.
	Table = "stadia"
	// StadiumConfirmationTable is the table the holds the StadiumConfirmation relation/edge.
	StadiumConfirmationTable = "confirmations"
	// StadiumConfirmationInverseTable is the table name for the Confirmation entity.
	// It exists in this package in order to avoid circular dependency with the "confirmation" package.
	StadiumConfirmationInverseTable = "confirmations"
	// StadiumConfirmationColumn is the table column denoting the StadiumConfirmation relation/edge.
	StadiumConfirmationColumn = "stadium_stadium_confirmation"
)

// Columns holds all SQL columns for stadium fields.
var Columns = []string{
	FieldID,
	FieldNamestadium,
}

var (
	// NamestadiumValidator is a validator for the "namestadium" field. It is called by the builders before save.
	NamestadiumValidator func(string) error
)
