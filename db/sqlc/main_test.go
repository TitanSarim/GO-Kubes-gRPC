package tutorial

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
    dbSource = "postgresql://postgres:12345@localhost:5432/simple_bank?sslmode=disable"
)

var testQuries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
    var err error
    testDB, err = pgxpool.New(context.Background(), dbSource)
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }
    defer testDB.Close()

    testQuries = New(testDB)

    os.Exit(m.Run())
}