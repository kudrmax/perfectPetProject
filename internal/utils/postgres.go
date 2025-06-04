package utils

import "strings"

func ReturningSQL(columns ...string) string {
	return "RETURNING " + strings.Join(columns, ", ")
}
