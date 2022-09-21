package testutil

import (
	"github.com/zhendong233/Books/pkg/bookserr"
)

var TestErr = bookserr.New(nil, bookserr.Unexpected, "test error") // Test Error
