package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	_ "github.com/jackc/pgx/v4/stdlib"
	dbxware "github.com/pocketbase/dbx"
)

type Dbx struct {
	helper.ORMInterface
	conn *dbxware.DB
}

func CreateDbx() helper.ORMInterface {
	return &Dbx{}
}

func (db *Dbx) Name() string {
	return "dbx"
}

func (db *Dbx) Init() error {
	dialect := "postgres"
	conn, err := dbxware.Open(dialect, helper.OrmSource)
	if err != nil {
		return err
	}

	db.conn = conn
	_, err = db.conn.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (db *Dbx) Close() error {
	return nil
}

func (db *Dbx) Insert(b *testing.B) {
	m := NewModel8()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.conn.Model(m).Insert()
		if err != nil {
			helper.SetError(b, db.Name(), "Insert", err.Error())
		}
	}
}

func (db *Dbx) InsertMulti(b *testing.B) {
	helper.SetError(b, db.Name(), "InsertMulti", "bulk-insert is not supported")
	//ms := make([]*Model, 0, 100)
	//for i := 0; i < 100; i++ {
	//	ms = append(ms, NewModel())
	//}
	//
	//b.ReportAllocs()
	//b.ResetTimer()
	//
	//for i := 0; i < b.N; i++ {
	//	for _, m := range ms {
	//		m.Id = 0
	//	}
	//
	//	_, err := gq.conn.Insert("models").
	//		Cols(columnsAny...).Rows(modelsToAnySlice(ms)...).Executor().Exec()
	//	if err != nil {
	//		helper.SetError(b, gq.Name(), "InsertMulti", err.Error())
	//	}
	//}
}

func (db *Dbx) Update(b *testing.B) {
	m := NewModel8()

	err := db.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, db.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.conn.Model(m).Update()
		if err != nil {
			helper.SetError(b, db.Name(), "Update", err.Error())
		}
	}
}

func (db *Dbx) Read(b *testing.B) {
	m := NewModel8()

	err := db.conn.Model(m).Insert()
	if err != nil {
		helper.SetError(b, db.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.conn.Select().Model(m.ID, m)
		if err != nil {
			helper.SetError(b, db.Name(), "Read", err.Error())
		}
	}
}

func (db *Dbx) ReadSlice(b *testing.B) {
	m := NewModel8()
	for i := 0; i < 100; i++ {
		err := db.conn.Model(m).Insert()
		if err != nil {
			helper.SetError(b, db.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var ms []*Model8
		err := db.conn.Select().From("models").
			Where(dbxware.NewExp("id>0")).Limit(100).All(&ms)
		if err != nil {
			helper.SetError(b, db.Name(), "ReadSlice", err.Error())
		}
	}
}
