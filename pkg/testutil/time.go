package testutil

import "time"

var (
	loc, _ = time.LoadLocation("Asia/Shanghai")
	// 以此纪念2020这糟心的一年
	TestTime, _ = time.ParseInLocation(time.RFC3339, "2020-01-01T08:00:00Z", loc)
)
