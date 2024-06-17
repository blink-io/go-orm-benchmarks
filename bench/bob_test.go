package bench

import (
	"github.com/efectn/go-orm-benchmarks/helper"
	"log"
	"runtime"
	"testing"
)

func TestBob(t *testing.T) {

}

func BenchmarkBob(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	helper.OrmSource = "user=postgres password=postgres host=localhost dbname=test sslmode=disable"

	// Clean tables for each run
	err := helper.CreateTables()
	if err != nil {
		panic(err)
	}

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
