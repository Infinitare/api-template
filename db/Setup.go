package db

import (
	"api-template/db/connection"
	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
)

func Setup() func() {

	str := os.Getenv("DATABASE_URL")
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithDBMPropagation(tracer.DBMPropagationModeFull), sqltrace.WithServiceName("supabase"))

	var err error
	db, err := sqltrace.Open("postgres", str)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(200)
	connection.Db = db

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return func() {
		_ = db.Close()
	}

}
