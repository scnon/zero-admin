package utils

import (
	"strconv"
	"strings"
)

func Int64sToStrings(ints []int64) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.FormatInt(v, 10)
	}
	return strs
}

func CreateDBPlaceholders(n int) string {
	if n <= 0 {
		return ""
	}
	placeholders := make([]string, n)
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}

func CreateJoinTableRows(name string, rows []string) string {
	if len(rows) == 0 {
		return ""
	}
	var rowsCopy = make([]string, len(rows))
	for i, row := range rows {
		rowsCopy[i] = name + "." + row
	}
	return strings.Join(rowsCopy, ", ")
}

func GetInt64ArryFromStr(str string) []int64 {
	strs := strings.Split(str, ",")
	var ids []int64
	for _, v := range strs {
		id, _ := strconv.ParseInt(v, 10, 64)
		ids = append(ids, id)
	}
	return ids
}
