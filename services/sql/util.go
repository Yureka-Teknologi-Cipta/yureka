package yureka_sql

import (
	"fmt"
	"strings"
)

// ConvertSQLBind converts SQL query bind from "?" to "$1", "$2", etc.
func ConvertSQLBind(query string) string {
	var i int
	return strings.ReplaceAll(query, "?", func() string {
		i++
		return fmt.Sprintf("$%d", i)
	}())
}
