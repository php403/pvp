package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_key").Unique(),
		field.String("app_screen").
			Default("unknown"),
		field.Int8("app_status").Default(1),
		field.Int16("create_time"),
		field.Int16("delete_time"),
		field.Int16("update_time"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}


