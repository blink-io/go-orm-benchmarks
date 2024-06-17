package bench

import (
	"github.com/aarondl/opt/omit"
	"github.com/efectn/go-orm-benchmarks/bench/bob/models"
	"github.com/efectn/go-orm-benchmarks/bench/bob/models/factory"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"

	_ "github.com/jackc/pgx/v4/stdlib"
	bobware "github.com/stephenafamo/bob"
)

type Bob struct {
	helper.ORMInterface
	conn bobware.DB
	f    *factory.Factory
}

func CreateBob() helper.ORMInterface {
	return &Bob{}
}

func (bob *Bob) Name() string {
	return "bob"
}

func (bob *Bob) Init() error {
	conn, err := bobware.Open("postgres", helper.OrmSource)
	if err != nil {
		return err
	}

	bob.conn = conn
	//bob.f = factory.New()
	//_, err = bob.conn.Begin()
	//if err != nil {
	//	return err
	//}

	return nil
}

func (bob *Bob) Close() error {
	return bob.conn.Close()
}

func (bob *Bob) Insert(b *testing.B) {
	m := NewModelAlt()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ms := newBobModelSetter(&m)
		_, err := models.Models.Insert(ctx, bob.conn, ms)
		if err != nil {
			helper.SetError(b, bob.Name(), "Insert", err.Error())
		}
	}
}

func newBobModelSetter(mat *Model) *models.ModelSetter {
	if mat == nil {
		m := NewModelAlt()
		mat = &m
	}
	ms := &models.ModelSetter{
		Name:    omit.From(mat.Name),
		Title:   omit.From(mat.Title),
		Fax:     omit.From(mat.Fax),
		Web:     omit.From(mat.Web),
		Age:     omit.From(int32(mat.Age)),
		Right:   omit.From(mat.Right),
		Counter: omit.From(mat.Counter),
	}
	return ms
}

func (bob *Bob) InsertMulti(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bulk := make([]*models.ModelSetter, 100)
	for i, _ := range bulk {
		bulk[i] = newBobModelSetter(nil)
	}

	for i := 0; i < b.N; i++ {
		_, err := models.Models.InsertMany(ctx, bob.conn, bulk...)
		if err != nil {
			helper.SetError(b, bob.Name(), "InsertMulti", err.Error())
		}
	}
}

func (bob *Bob) Update(b *testing.B) {
	ms := newBobModelSetter(nil)

	mr, err := models.Models.Insert(ctx, bob.conn, ms)
	if err != nil {
		helper.SetError(b, bob.Name(), "Update", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ums := &models.ModelSetter{
			Name:    ms.Name,
			Title:   ms.Title,
			Fax:     ms.Fax,
			Web:     ms.Web,
			Age:     ms.Age,
			Right:   ms.Right,
			Counter: ms.Counter,
		}
		err := models.Models.Update(ctx, bob.conn, ums, mr)
		if err != nil {
			helper.SetError(b, bob.Name(), "Update", err.Error())
		}
	}
}

func (bob *Bob) Read(b *testing.B) {
	ms := newBobModelSetter(nil)

	mr, err := models.Models.Insert(ctx, bob.conn, ms)
	if err != nil {
		helper.SetError(b, bob.Name(), "Read", err.Error())
	}

	b.ReportAllocs()
	b.ResetTimer()

	wq := models.SelectWhere.Models.ID.EQ(mr.ID)

	for i := 0; i < b.N; i++ {
		//_, err := models.FindModel(ctx, bob.conn, mr.ID)
		_, err := models.Models.Query(ctx, bob.conn, wq).One()
		if err != nil {
			helper.SetError(b, bob.Name(), "Read", err.Error())
		}
	}
}

func (bob *Bob) ReadSlice(b *testing.B) {
	ms := newBobModelSetter(nil)
	for i := 0; i < 100; i++ {
		_, err := models.Models.Insert(ctx, bob.conn, ms)
		if err != nil {
			helper.SetError(b, bob.Name(), "ReadSlice", err.Error())
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	wq := models.SelectWhere.Models.ID.GT(0)

	for i := 0; i < b.N; i++ {
		//_, err := bob.conn.Select("*").From("models").Where("id > 0").Limit(100).Load(&ms)
		_, err := models.Models.Query(ctx, bob.conn, wq, sm.Limit(100)).All()
		if err != nil {
			helper.SetError(b, bob.Name(), "ReadSlice", err.Error())
		}
	}
}
