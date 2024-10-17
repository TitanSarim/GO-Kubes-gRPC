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

func TestMain(m *testing.M){
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil{
		log.Fatal("Cannot connect to database")
	}
	defer conn.Close()

	testQuries = New(conn)

	os.Exit(m.Run())
}