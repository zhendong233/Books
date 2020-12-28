package dbutil

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
)

var (
	Conn    = fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", User, Pass, Port, DB)
	ConnStr = Conn + `?charset=utf8mb4&parseTime=True&loc=UTC&tls=false&multiStatements=true`
)

func NewDB() (*sql.DB, error) {
	d, err := sql.Open("mysql", ConnStr)
	if err != nil {
		return nil, err
	}
	return d, nil
}
