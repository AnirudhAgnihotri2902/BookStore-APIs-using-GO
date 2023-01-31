// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/books"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Books is the model entity for the Books schema.
type Books struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BookName holds the value of the "book_name" field.
	BookName string `json:"book_name,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// PublishDate holds the value of the "publish_date" field.
	PublishDate string `json:"publish_date,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Books) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case books.FieldID:
			values[i] = new(sql.NullInt64)
		case books.FieldBookName, books.FieldAuthor, books.FieldPublishDate:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Books", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Books fields.
func (b *Books) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case books.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case books.FieldBookName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field book_name", values[i])
			} else if value.Valid {
				b.BookName = value.String
			}
		case books.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				b.Author = value.String
			}
		case books.FieldPublishDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publish_date", values[i])
			} else if value.Valid {
				b.PublishDate = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Books.
// Note that you need to call Books.Unwrap() before calling this method if this Books
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Books) Update() *BooksUpdateOne {
	return NewBooksClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Books entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Books) Unwrap() *Books {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Books is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Books) String() string {
	var builder strings.Builder
	builder.WriteString("Books(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("book_name=")
	builder.WriteString(b.BookName)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(b.Author)
	builder.WriteString(", ")
	builder.WriteString("publish_date=")
	builder.WriteString(b.PublishDate)
	builder.WriteByte(')')
	return builder.String()
}

// BooksSlice is a parsable slice of Books.
type BooksSlice []*Books

func (b BooksSlice) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}