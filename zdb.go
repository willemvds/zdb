package zdb

import (
	"database/sql"
	"errors"
)

var ErrNoDb = errors.New("zdb: no database configured")

type ZDB struct {
	*sql.DB
}

func (zdb* ZDB) Set(db *sql.DB) {
	zdb.DB = db
}

func (zdb *ZDB) Ping() {
	if zdb.DB == nil {
		return
	}

	zdb.DB.Ping()
}

func (zdb *ZDB) Close() {
	if zdb.DB == nil {
		return
	}
	zdb.DB.Close()	
}

func (zdb *ZDB) SetMaxIdleConns(n int) {
	if zdb.DB == nil {
		return
	}
	zdb.DB.SetMaxIdleConns(n)
}

func (zdb *ZDB) SetMaxOpenConns(n int) {
	if zdb.DB == nil {
		return
	}
	zdb.DB.SetMaxOpenConns(n)
}

func (zdb *ZDB) Stats() sql.DBStats {
	if zdb.DB == nil {
		return sql.DBStats{}
	}
	return zdb.DB.Stats()
}

func (zdb *ZDB) Prepare(query string) (*sql.Stmt, error) {
	if zdb.DB == nil {
		return nil, ErrNoDb
	}
	return zdb.DB.Prepare(query)
}

func (zdb *ZDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if zdb.DB == nil {
		return nil, ErrNoDb
	}
	return zdb.DB.Exec(query, args...)
}

func (zdb *ZDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if zdb.DB == nil {
		return nil, ErrNoDb
	}
	return zdb.DB.Query(query, args...)
}

func (zdb *ZDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if zdb.DB == nil {
		return &sql.Row{}
	}
	return zdb.DB.QueryRow(query, args...)
}

func (zdb *ZDB) Begin() (*sql.Tx, error) {
	if zdb.DB == nil {
		return nil, ErrNoDb
	}
	return zdb.DB.Begin()
}
