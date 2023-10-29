package utils

import (
	"fmt"
	"strings"
)

func ConvertPlaceholdersMySqlToPostgreSQL(query string) string {
	var convertedQuery strings.Builder
	var placeholderCount int

	for _, char := range query {
		if char == '?' {
			placeholderCount++
			convertedQuery.WriteString(fmt.Sprintf("$%d", placeholderCount))
		} else {
			convertedQuery.WriteRune(char)
		}
	}

	return convertedQuery.String()
}
