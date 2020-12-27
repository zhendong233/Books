package dbutil

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	Conn    = fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", User, Pass, Port, DB)
	ConnStr = Conn + `?charset=utf8mb4&parseTime=True&loc=UTC&tls=false&multiStatements=true`
)

func NewDB() *sql.DB {
	d, err := sql.Open("mysql", ConnStr)
	if err != nil {
		os.Exit(-1)
	}
	return d
}
