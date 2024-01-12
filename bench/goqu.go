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

func (gq *Goqu) Name() string {
	return "goqu"
}

func (gq *Goqu) Init() error {
	dialect := "postgres"
	conn, err := sql.Open(dialect, helper.OrmSource)
	if err != nil {
		return err
	}

	gq.conn = goquware.New(dialect, conn)
	_, err = gq.conn.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (gq *Goqu) Close() error {
	return nil
}

func (gq *Goqu) Insert(b *testing.B) {
	m := NewModel()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gq.conn.Insert("models").
			Cols(columnsAny...).Rows(m).Executor().Exec()
		if err != nil {
			helper.SetError(b, gq.Name(), "Insert", err.Error())
		}
	}
}

func modelsToAnySlice(ms []*Model) []any {
	as := make([]any, len(ms))
	for i, m := range ms {
		as[i] = m
	}
	return as
}

func (gq *Goqu) InsertMulti(b *testing.B) {
	ms := make([]*Model, 0, 100)
	for i := 0; i < 100; i++ {
		ms = append(ms, NewModel())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}

		_, err := gq.conn.Insert("models").
			Cols(columnsAny...).Rows(modelsToAnySlice(ms)...).Executor().Exec()
		if err != nil {
			helper.SetError(b, gq.Name(), "InsertMulti", err.Error())
		}
	}
}

func (gq *Goqu) Update(b *testing.B) {
	m := NewModel()

	_, err := gq.conn.Insert("models").
		Cols(columnsAny...).Rows(m).Executor().Exec()
	if err != nil {
		helper.SetError(b, gq.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gq.conn.Update("models").Set(map[string]any{
			"name":    m.Name,
			"title":   m.Title,
			"fax":     m.Fax,
			"web":     m.Web,
			"age":     m.Age,
			"right":   m.Right,
			"counter": m.Counter,
		}).Executor().Exec()
		if err != nil {
			helper.SetError(b, gq.Name(), "Update", err.Error())
		}
	}
}

func (gq *Goqu) Read(b *testing.B) {
	m := NewModel()

	_, err := gq.conn.Insert("models").Cols(columnsAny...).Rows(m).Executor().Exec()
	if err != nil {
		helper.SetError(b, gq.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gq.conn.From("models").Select(goquware.Star()).
			Where(goquware.C("id").Eq(m.Id)).ScanStruct(m)
		if err != nil {
			helper.SetError(b, gq.Name(), "Read", err.Error())
		}
	}
}

func (gq *Goqu) ReadSlice(b *testing.B) {
	m := NewModel()
	for i := 0; i < 100; i++ {
		_, err := gq.conn.Insert("models").Cols(columnsAny...).Rows(m).Executor().Exec()
		if err != nil {
			helper.SetError(b, gq.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model2
		err := gq.conn.From("models").
			Select(goquware.Star()).
			Where(goquware.C("id").Gt(0)).Limit(100).ScanStructs(&ms)
		if err != nil {
			helper.SetError(b, gq.Name(), "ReadSlice", err.Error())
		}
	}
}
