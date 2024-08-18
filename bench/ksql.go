package bench

import (
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

var ModelTable = ksql.NewTable("models")

type Ksql struct {
	helper.ORMInterface
	conn ksql.DB
}

func CreateKsql() helper.ORMInterface {
	return &Ksql{}
}

func (v *Ksql) Name() string {
	return "ksql"
}

func (v *Ksql) Init() error {
	conn, err := kpgx.New(ctx, helper.OrmSource, ksql.Config{
		MaxOpenConns: helper.OrmMaxConn,
	})
	if err != nil {
		return err
	}

	v.conn = conn

	return nil
}

func (v *Ksql) Close() error {
	return v.conn.Close()
}

func (v *Ksql) Insert(b *testing.B) {
	m := NewModel9()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := v.conn.Insert(ctx, ModelTable, m)
		if err != nil {
			helper.SetError(b, v.Name(), "Insert", err.Error())
		}
	}
}

func (v *Ksql) InsertMulti(b *testing.B) {
	helper.SetError(b, v.Name(), "InsertMulti", "bulk-insert is not supported")
}

func (v *Ksql) Update(b *testing.B) {
	ms := NewModel9()

	err := v.conn.Insert(ctx, ModelTable, ms)
	if err != nil {
		helper.SetError(b, v.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := v.conn.Patch(ctx, ModelTable, &Model9{
			ID:      ms.ID,
			Name:    ms.Name,
			Title:   ms.Title,
			Fax:     ms.Fax,
			Web:     ms.Web,
			Age:     ms.Age,
			Right:   ms.Right,
			Counter: ms.Counter,
		})
		if err != nil {
			helper.SetError(b, v.Name(), "Update", err.Error())
		}
	}
}

func (v *Ksql) Read(b *testing.B) {
	ms := NewModel9()

	err := v.conn.Insert(ctx, ModelTable, ms)
	if err != nil {
		helper.SetError(b, v.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var rm Model9
		err := v.conn.QueryOne(ctx, &rm, "select * from models where id = $1", ms.ID)
		if err != nil {
			helper.SetError(b, v.Name(), "Read", err.Error())
		}
	}
}

func (v *Ksql) ReadSlice(b *testing.B) {
	ms := NewModel9()
	for i := 0; i < 100; i++ {
		ms.ID = 0
		err := v.conn.Insert(ctx, ModelTable, ms)
		if err != nil {
			helper.SetError(b, v.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var maa []Model9
		err := v.conn.Query(ctx, &maa, "select * from models where id >0 limit 100")
		if err != nil {
			helper.SetError(b, v.Name(), "ReadSlice", err.Error())
		}
	}
}
