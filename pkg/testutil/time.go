package testutil

import "time"

var TestTime, _ = time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")  // 以此纪念2020这糟心的一年
var ATestTime, _ = time.Parse(time.RFC3339, "2020-01-01T16:00:00Z") // 本地总是差8小时
