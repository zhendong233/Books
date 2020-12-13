package dbutil

import (
	"bufio"
	"context"
	"database/sql"
	"os"
	"regexp"
	"strings"
)

func ExecSQLFile(ctx context.Context, db *sql.DB, filePath string) error {
	f, err := os.Open(filePath)
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
			if _, err := db.ExecContext(ctx, q); err != nil {
				return err
			}
		}
	}
	return nil
}
