package testutil

import (
	"context"
	"testing"
	"time"
)

func Test_SetFakeTimeForMysql(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db := PrepareMySQL(t)
	defer db.Close()

	SetFakeTimeForMysql(t, db, TestTime)
	var actual time.Time
	//创建一个默认的ctx
	ctx := context.Background()
	err := db.QueryRowContext(ctx, "SELECT NOW()").Scan(&actual)
	if err != nil {
		t.Errorf("select error %s", err)
	}
	if err == nil && actual != TestTime {
		t.Error("failure at set TIMESTAMP")
	}
}
