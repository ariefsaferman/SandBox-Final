package utils

import (
	"strings"
)

func SplitStrToSlice(s string) []string {
	return strings.Split(s, " ")
}
