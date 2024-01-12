package bench

import (
	"fmt"
	"os"
	"testing"

	"github.com/efectn/go-orm-benchmarks/helper"
	dbxware "github.com/pocketbase/dbx"
	"github.com/stretchr/testify/require"
)

func TestDbx_Insert(t *testing.T) {
	pgDSN := os.Getenv("PG_DSN")
	fmt.Println("DSN: ", pgDSN)

	require.NotEmpty(t, pgDSN)

	helper.OrmSource = pgDSN
	dbx := &Dbx{}
	require.NoError(t, dbx.Init())

	r1 := NewModel8()

	for i := 0; i < 10; i++ {
		err := dbx.conn.Model(r1).Insert()
		require.NoError(t, err)
		r1.ID = 0
	}
}

func TestSelect(t *testing.T) {
	pgDSN := os.Getenv("PG_DSN")
	fmt.Println("DSN: ", pgDSN)

	require.NotEmpty(t, pgDSN)

	helper.OrmSource = pgDSN
	dbx := &Dbx{}
	require.NoError(t, dbx.Init())

	var ms []*Model8

	err := dbx.conn.Select().From("models").
		Where(dbxware.NewExp("id>0")).Limit(100).All(&ms)

	require.NoError(t, err)
}
