package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName string, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatalf("db connection err %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("db ping err %v", err)
	}

	return &Adapter{
		db: db,
	}, nil
}

func (da *Adapter) CloseDBConnection() {
	err := da.db.Close()

	if err != nil {
		log.Fatalf("")
	}
}

func (da *Adapter) AddToHistory(ans int32, operation string) error {
	querryString, args, err := sq.Insert("arith_history").Columns("date", "answere", "operation").Values(time.Now(), ans, operation).ToSql()
	if err != nil {
		return err
	}

	if _, err := da.db.Exec(querryString, args...); err != nil {
		return err
	}
	return nil
}
