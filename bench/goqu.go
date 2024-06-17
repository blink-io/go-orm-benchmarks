package bench

import (
	"database/sql"
	"testing"

	goquware "github.com/doug-martin/goqu/v9"
	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var columnsAny = []any{"name", "title", "fax", "web", "age", "right", "counter"}

type Goqu struct {
	helper.ORMInterface
	conn *goquware.Database
}

func CreateGoqu() helper.ORMInterface {
	return &Goqu{}
}

func (db *Goqu) Name() string {
	return "goqu"
}

func (db *Goqu) Init() error {
	dialect := "postgres"
	conn, err := sql.Open(dialect, helper.OrmSource)
	if err != nil {
		return err
	}

	db.conn = goquware.New(dialect, conn)
	_, err = db.conn.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (db *Goqu) Close() error {
	return nil
}

func (db *Goqu) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.conn.Insert("models").
			Cols(columnsAny...).Rows(m).Executor().Exec()
		if err != nil {
			helper.SetError(b, db.Name(), "Insert", err.Error())
		}
	}
}

func (db *Goqu) InsertMulti(b *testing.B) {
	ms := make([]any, 0, 100)
	for i := 0; i < 100; i++ {
		mi := NewModel()
		mi.Id = 0
		ms = append(ms, mi)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.conn.Insert("models").
			Cols(columnsAny...).Rows(ms...).Executor().Exec()
		if err != nil {
			helper.SetError(b, db.Name(), "InsertMulti", err.Error())
		}
	}
}

func (db *Goqu) Update(b *testing.B) {
	m := NewModel()

	_, err := db.conn.Insert("models").
		Cols(columnsAny...).Rows(m).Executor().Exec()
	if err != nil {
		helper.SetError(b, db.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.conn.Update("models").Set(map[string]any{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Executor().Exec()
		if err != nil {
			helper.SetError(b, db.Name(), "Update", err.Error())
		}
	}
}

func (db *Goqu) Read(b *testing.B) {
	m := NewModel()

	_, err := db.conn.Insert("models").Cols(columnsAny...).Rows(m).Executor().Exec()
	if err != nil {
		helper.SetError(b, db.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.conn.From("models").Select(goquware.Star()).
			Where(goquware.C("id").Eq(m.Id)).ScanStruct(m)
		if err != nil {
			helper.SetError(b, db.Name(), "Read", err.Error())
		}
	}
}

func (db *Goqu) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := db.conn.Insert("models").Cols(columnsAny...).Rows(m).Executor().Exec()
		if err != nil {
			helper.SetError(b, db.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		err := db.conn.From("models").
			Select(goquware.Star()).
			Where(goquware.C("id").Gt(0)).Limit(100).ScanStructs(&ms)
		if err != nil {
			helper.SetError(b, db.Name(), "ReadSlice", err.Error())
		}
	}
}
