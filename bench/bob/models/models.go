// Code generated by BobGen psql v0.27.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/bob/expr"
)

// Model is an object representing the database table.
type Model struct {
	ID      int32  `db:"id,pk" json:"id" toml:"id" yaml:"id" msgpack:"id" bun:"id"`
	Name    string `db:"name" json:"name" toml:"name" yaml:"name" msgpack:"name" bun:"name"`
	Title   string `db:"title" json:"title" toml:"title" yaml:"title" msgpack:"title" bun:"title"`
	Fax     string `db:"fax" json:"fax" toml:"fax" yaml:"fax" msgpack:"fax" bun:"fax"`
	Web     string `db:"web" json:"web" toml:"web" yaml:"web" msgpack:"web" bun:"web"`
	Age     int32  `db:"age" json:"age" toml:"age" yaml:"age" msgpack:"age" bun:"age"`
	Right   bool   `db:"right" json:"right" toml:"right" yaml:"right" msgpack:"right" bun:"right"`
	Counter int64  `db:"counter" json:"counter" toml:"counter" yaml:"counter" msgpack:"counter" bun:"counter"`
}

// ModelSlice is an alias for a slice of pointers to Model.
// This should almost always be used instead of []*Model.
type ModelSlice []*Model

// Models contains methods to work with the models table
var Models = psql.NewTablex[*Model, ModelSlice, *ModelSetter]("", "models")

// ModelsQuery is a query on the models table
type ModelsQuery = *psql.ViewQuery[*Model, ModelSlice]

// ModelsStmt is a prepared statment on models
type ModelsStmt = bob.QueryStmt[*Model, ModelSlice]

// ModelSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type ModelSetter struct {
	ID      omit.Val[int32]  `db:"id,pk" json:"id" toml:"id" yaml:"id" msgpack:"id" bun:"id"`
	Name    omit.Val[string] `db:"name" json:"name" toml:"name" yaml:"name" msgpack:"name" bun:"name"`
	Title   omit.Val[string] `db:"title" json:"title" toml:"title" yaml:"title" msgpack:"title" bun:"title"`
	Fax     omit.Val[string] `db:"fax" json:"fax" toml:"fax" yaml:"fax" msgpack:"fax" bun:"fax"`
	Web     omit.Val[string] `db:"web" json:"web" toml:"web" yaml:"web" msgpack:"web" bun:"web"`
	Age     omit.Val[int32]  `db:"age" json:"age" toml:"age" yaml:"age" msgpack:"age" bun:"age"`
	Right   omit.Val[bool]   `db:"right" json:"right" toml:"right" yaml:"right" msgpack:"right" bun:"right"`
	Counter omit.Val[int64]  `db:"counter" json:"counter" toml:"counter" yaml:"counter" msgpack:"counter" bun:"counter"`
}

func (s ModelSetter) SetColumns() []string {
	vals := make([]string, 0, 8)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.Name.IsUnset() {
		vals = append(vals, "name")
	}

	if !s.Title.IsUnset() {
		vals = append(vals, "title")
	}

	if !s.Fax.IsUnset() {
		vals = append(vals, "fax")
	}

	if !s.Web.IsUnset() {
		vals = append(vals, "web")
	}

	if !s.Age.IsUnset() {
		vals = append(vals, "age")
	}

	if !s.Right.IsUnset() {
		vals = append(vals, "right")
	}

	if !s.Counter.IsUnset() {
		vals = append(vals, "counter")
	}

	return vals
}

func (s ModelSetter) Overwrite(t *Model) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.Name.IsUnset() {
		t.Name, _ = s.Name.Get()
	}
	if !s.Title.IsUnset() {
		t.Title, _ = s.Title.Get()
	}
	if !s.Fax.IsUnset() {
		t.Fax, _ = s.Fax.Get()
	}
	if !s.Web.IsUnset() {
		t.Web, _ = s.Web.Get()
	}
	if !s.Age.IsUnset() {
		t.Age, _ = s.Age.Get()
	}
	if !s.Right.IsUnset() {
		t.Right, _ = s.Right.Get()
	}
	if !s.Counter.IsUnset() {
		t.Counter, _ = s.Counter.Get()
	}
}

func (s ModelSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 8)
	if s.ID.IsUnset() {
		vals[0] = psql.Raw("DEFAULT")
	} else {
		vals[0] = psql.Arg(s.ID)
	}

	if s.Name.IsUnset() {
		vals[1] = psql.Raw("DEFAULT")
	} else {
		vals[1] = psql.Arg(s.Name)
	}

	if s.Title.IsUnset() {
		vals[2] = psql.Raw("DEFAULT")
	} else {
		vals[2] = psql.Arg(s.Title)
	}

	if s.Fax.IsUnset() {
		vals[3] = psql.Raw("DEFAULT")
	} else {
		vals[3] = psql.Arg(s.Fax)
	}

	if s.Web.IsUnset() {
		vals[4] = psql.Raw("DEFAULT")
	} else {
		vals[4] = psql.Arg(s.Web)
	}

	if s.Age.IsUnset() {
		vals[5] = psql.Raw("DEFAULT")
	} else {
		vals[5] = psql.Arg(s.Age)
	}

	if s.Right.IsUnset() {
		vals[6] = psql.Raw("DEFAULT")
	} else {
		vals[6] = psql.Arg(s.Right)
	}

	if s.Counter.IsUnset() {
		vals[7] = psql.Raw("DEFAULT")
	} else {
		vals[7] = psql.Arg(s.Counter)
	}

	return im.Values(vals...)
}

func (s ModelSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s ModelSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 8)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "id")...),
			psql.Arg(s.ID),
		}})
	}

	if !s.Name.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "name")...),
			psql.Arg(s.Name),
		}})
	}

	if !s.Title.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "title")...),
			psql.Arg(s.Title),
		}})
	}

	if !s.Fax.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "fax")...),
			psql.Arg(s.Fax),
		}})
	}

	if !s.Web.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "web")...),
			psql.Arg(s.Web),
		}})
	}

	if !s.Age.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "age")...),
			psql.Arg(s.Age),
		}})
	}

	if !s.Right.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "right")...),
			psql.Arg(s.Right),
		}})
	}

	if !s.Counter.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "counter")...),
			psql.Arg(s.Counter),
		}})
	}

	return exprs
}

type modelColumnNames struct {
	ID      string
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     string
	Right   string
	Counter string
}

var ModelColumns = buildModelColumns("models")

type modelColumns struct {
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

func (c modelColumns) Alias() string {
	return c.tableAlias
}

func (modelColumns) AliasedAs(alias string) modelColumns {
	return buildModelColumns(alias)
}

func buildModelColumns(alias string) modelColumns {
	return modelColumns{
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

type modelWhere[Q psql.Filterable] struct {
	ID      psql.WhereMod[Q, int32]
	Name    psql.WhereMod[Q, string]
	Title   psql.WhereMod[Q, string]
	Fax     psql.WhereMod[Q, string]
	Web     psql.WhereMod[Q, string]
	Age     psql.WhereMod[Q, int32]
	Right   psql.WhereMod[Q, bool]
	Counter psql.WhereMod[Q, int64]
}

func (modelWhere[Q]) AliasedAs(alias string) modelWhere[Q] {
	return buildModelWhere[Q](buildModelColumns(alias))
}

func buildModelWhere[Q psql.Filterable](cols modelColumns) modelWhere[Q] {
	return modelWhere[Q]{
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

// FindModel retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindModel(ctx context.Context, exec bob.Executor, IDPK int32, cols ...string) (*Model, error) {
	if len(cols) == 0 {
		return Models.Query(
			ctx, exec,
			SelectWhere.Models.ID.EQ(IDPK),
		).One()
	}

	return Models.Query(
		ctx, exec,
		SelectWhere.Models.ID.EQ(IDPK),
		sm.Columns(Models.Columns().Only(cols...)),
	).One()
}

// ModelExists checks the presence of a single record by primary key
func ModelExists(ctx context.Context, exec bob.Executor, IDPK int32) (bool, error) {
	return Models.Query(
		ctx, exec,
		SelectWhere.Models.ID.EQ(IDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the Model
func (o *Model) PrimaryKeyVals() bob.Expression {
	return psql.Arg(o.ID)
}

// Update uses an executor to update the Model
func (o *Model) Update(ctx context.Context, exec bob.Executor, s *ModelSetter) error {
	return Models.Update(ctx, exec, s, o)
}

// Delete deletes a single Model record with an executor
func (o *Model) Delete(ctx context.Context, exec bob.Executor) error {
	return Models.Delete(ctx, exec, o)
}

// Reload refreshes the Model using the executor
func (o *Model) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := Models.Query(
		ctx, exec,
		SelectWhere.Models.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}

	*o = *o2

	return nil
}

func (o ModelSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals ModelSetter) error {
	return Models.Update(ctx, exec, &vals, o...)
}

func (o ModelSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return Models.Delete(ctx, exec, o...)
}

func (o ModelSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]int32, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.Models.ID.In(IDPK...),
	)

	o2, err := Models.Query(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}

			*old = *new
			break
		}
	}

	return nil
}
