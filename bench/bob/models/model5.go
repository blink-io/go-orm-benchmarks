// Code generated by BobGen psql v0.27.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
)

// Model5 is an object representing the database table.
type Model5 struct {
	ID      int32  `db:"id" json:"id" toml:"id" yaml:"id" msgpack:"id" bun:"id"`
	Name    string `db:"name" json:"name" toml:"name" yaml:"name" msgpack:"name" bun:"name"`
	Title   string `db:"title" json:"title" toml:"title" yaml:"title" msgpack:"title" bun:"title"`
	Fax     string `db:"fax" json:"fax" toml:"fax" yaml:"fax" msgpack:"fax" bun:"fax"`
	Web     string `db:"web" json:"web" toml:"web" yaml:"web" msgpack:"web" bun:"web"`
	Age     int32  `db:"age" json:"age" toml:"age" yaml:"age" msgpack:"age" bun:"age"`
	Right   bool   `db:"right" json:"right" toml:"right" yaml:"right" msgpack:"right" bun:"right"`
	Counter int64  `db:"counter" json:"counter" toml:"counter" yaml:"counter" msgpack:"counter" bun:"counter"`
}

// Model5Slice is an alias for a slice of pointers to Model5.
// This should almost always be used instead of []*Model5.
type Model5Slice []*Model5

// Model5s contains methods to work with the model5 view
var Model5s = psql.NewViewx[*Model5, Model5Slice]("", "model5")

// Model5sQuery is a query on the model5 view
type Model5sQuery = *psql.ViewQuery[*Model5, Model5Slice]

// Model5sStmt is a prepared statment on model5
type Model5sStmt = bob.QueryStmt[*Model5, Model5Slice]

type model5ColumnNames struct {
	ID      string
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     string
	Right   string
	Counter string
}

var Model5Columns = buildModel5Columns("model5")

type model5Columns struct {
	tableAlias string
	ID         psql.Expression
	Name       psql.Expression
	Title      psql.Expression
	Fax        psql.Expression
	Web        psql.Expression
	Age        psql.Expression
	Right      psql.Expression
	Counter    psql.Expression
}

func (c model5Columns) Alias() string {
	return c.tableAlias
}

func (model5Columns) AliasedAs(alias string) model5Columns {
	return buildModel5Columns(alias)
}

func buildModel5Columns(alias string) model5Columns {
	return model5Columns{
		tableAlias: alias,
		ID:         psql.Quote(alias, "id"),
		Name:       psql.Quote(alias, "name"),
		Title:      psql.Quote(alias, "title"),
		Fax:        psql.Quote(alias, "fax"),
		Web:        psql.Quote(alias, "web"),
		Age:        psql.Quote(alias, "age"),
		Right:      psql.Quote(alias, "right"),
		Counter:    psql.Quote(alias, "counter"),
	}
}

type model5Where[Q psql.Filterable] struct {
	ID      psql.WhereMod[Q, int32]
	Name    psql.WhereMod[Q, string]
	Title   psql.WhereMod[Q, string]
	Fax     psql.WhereMod[Q, string]
	Web     psql.WhereMod[Q, string]
	Age     psql.WhereMod[Q, int32]
	Right   psql.WhereMod[Q, bool]
	Counter psql.WhereMod[Q, int64]
}

func (model5Where[Q]) AliasedAs(alias string) model5Where[Q] {
	return buildModel5Where[Q](buildModel5Columns(alias))
}

func buildModel5Where[Q psql.Filterable](cols model5Columns) model5Where[Q] {
	return model5Where[Q]{
		ID:      psql.Where[Q, int32](cols.ID),
		Name:    psql.Where[Q, string](cols.Name),
		Title:   psql.Where[Q, string](cols.Title),
		Fax:     psql.Where[Q, string](cols.Fax),
		Web:     psql.Where[Q, string](cols.Web),
		Age:     psql.Where[Q, int32](cols.Age),
		Right:   psql.Where[Q, bool](cols.Right),
		Counter: psql.Where[Q, int64](cols.Counter),
	}
}
