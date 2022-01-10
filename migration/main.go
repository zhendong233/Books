package main

import (
	"context"
	"log"
	"os"

	"github.com/zhendong233/Books/pkg/dbutil"
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}

func run() error {
	db, err := dbutil.NewDB()
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	if err := dbutil.ExecSQLFile(context.Background(), db, "./migration/sql/createTable.sql"); err != nil {
		return err
	}
	return nil
}
