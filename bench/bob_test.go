package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"log"
	"testing"
)

func TestBob(t *testing.T) {

}

func BenchmarkBob(b *testing.B) {
	helper.OrmSource = "user=postgres password=postgres dbname=test sslmode=disable"
	bb := CreateBob()
	if err := bb.Init(); err != nil {
		log.Fatal("init bob fail")
	}
	bb.Insert(b)
	bb.InsertMulti(b)
	bb.Update(b)
	bb.Read(b)
	bb.ReadSlice(b)

	//
	_ = bb.Close()
}
