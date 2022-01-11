package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql

	"github.com/zhendong233/Books/pkg/dbutil"
)

const (
	Conn    = "books:books@tcp(localhost:3309)/books"
	ConnStr = Conn + `?charset=utf8mb4&parseTime=True&loc=UTC&tls=false&multiStatements=true`
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}

func run() error {
	db, err := sql.Open("mysql", ConnStr)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	db.SetConnMaxLifetime(5)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	if err := dbutil.ExecSQLFile(context.Background(), db, "./migration/sql/createTable.sql"); err != nil {
		return err
	}
	return nil
}
