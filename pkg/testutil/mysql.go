package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql" // mysql

	"github.com/zhendong233/Books/pkg/dbutil"
)

const (
	TestBookID = "eae9827c-349f-43a2-82c1-0ef863cfd5ba"
)

var (
	driverName  = "mysql"
	TestConn    = fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", dbutil.User, dbutil.Pass, dbutil.Port, dbutil.DB)
	TestConnStr = TestConn + `?charset=utf8mb4&parseTime=True&loc=UTC&tls=false&multiStatements=true`
	mu          sync.Mutex
)

var txDBRegisterOnce sync.Once

func registerDB() {
	txdb.Register("test-mysql", driverName, TestConnStr)
}

func PrepareMySQL(t *testing.T) *sql.DB {
	t.Helper()
	txDBRegisterOnce.Do(registerDB)
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano()) // 时间戳（纳秒) 毫秒的话可以使用UnixNano() / 1e6
	db, err := sql.Open("test-mysql", cName)
	if err != nil {
		t.Fatalf("failed to open test-mysql connection: %s", err)
	}
	return db
}

func SetFakeTimeForMysql(t *testing.T, db *sql.DB, fakeTime time.Time) {
	t.Helper()
	if _, err := db.Exec("SET TIMESTAMP = ?", fakeTime.Unix()); err != nil {
		t.Fatal(err)
	}
}

func ExecSQLFile(t *testing.T, db *sql.DB, filePaths ...string) {
	t.Helper()
	mu.Lock()
	defer mu.Unlock()
	ctx := context.Background()
	for _, path := range filePaths {
		if err := dbutil.ExecSQLFile(ctx, db, path); err != nil {
			t.Fatal(err)
		}
	}
}
