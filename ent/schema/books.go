package schema

import(
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)


// Books holds the schema definition for the Books entity.
type Books struct {
	ent.Schema
}

// Fields of the Books.
func (Books) Fields() []ent.Field {
	return []ent.Field{
			field.String("book_name"),
			field.String("author").
					Default("Anonymous"),
			field.String("publish_date"),
	}
}

// Edges of the Books.
func (Books) Edges() []ent.Edge {
	return nil
}
