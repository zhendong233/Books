package testutil

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewSQLMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
	})
	return db, mock
}

func NewSQLMockDB(t *testing.T) *sql.DB {
	t.Helper()
	db, _ := NewSQLMock(t)
	return db
}
