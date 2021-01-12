// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "app_key", Type: field.TypeString, Unique: true},
		{Name: "app_screen", Type: field.TypeString, Default: "unknown"},
		{Name: "app_status", Type: field.TypeInt8, Default: 1},
		{Name: "create_time", Type: field.TypeInt16},
		{Name: "delete_time", Type: field.TypeInt16},
		{Name: "update_time", Type: field.TypeInt16},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:        "apps",
		Columns:     AppsColumns,
		PrimaryKey:  []*schema.Column{AppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
	}
)

func init() {
}
