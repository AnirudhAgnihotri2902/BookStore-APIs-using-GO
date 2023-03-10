// Code generated by ent, DO NOT EDIT.

package books

const (
	// Label holds the string label denoting the books type in the database.
	Label = "books"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBookName holds the string denoting the book_name field in the database.
	FieldBookName = "book_name"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldPublishDate holds the string denoting the publish_date field in the database.
	FieldPublishDate = "publish_date"
	// Table holds the table name of the books in the database.
	Table = "books"
)

// Columns holds all SQL columns for books fields.
var Columns = []string{
	FieldID,
	FieldBookName,
	FieldAuthor,
	FieldPublishDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAuthor holds the default value on creation for the "author" field.
	DefaultAuthor string
)
