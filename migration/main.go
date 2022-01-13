package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql
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
	f, err := os.Open("./migration/sql/createTable.sql")
	if err != nil {
		return err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	var sb strings.Builder
	reg := regexp.MustCompile(`--.*$`)
	for scan.Scan() {
		line := scan.Text()
		q := reg.ReplaceAllString(line, "") // 去掉sql文件中的注释
		sb.WriteString(q)
	}
	qs := strings.Split(sb.String(), ";")
	for _, q := range qs {
		if q != "" {
			if _, err := db.Exec(q); err != nil {
				return err
			}
		}
	}
	return nil
}
