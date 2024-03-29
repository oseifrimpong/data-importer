package utils

import (
	"regexp"
	"strings"
)

// SortOrder returns the string for sorting and orderin data
func SortOrder(table, sort, order string) string {
	return table + "." + toSnakeCase(sort) + " " + toSnakeCase(order)
}

// toSnakeCase changes string to database table
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
