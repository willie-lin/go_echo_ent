// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// Table holds the table name of the group in the database.
	Table = "groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Group type.
var ForeignKeys = []string{
	"user_groups",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
