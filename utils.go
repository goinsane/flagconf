package flagconf

import (
	"strconv"
)

func tryUnquote(q string) string {
	result, err := strconv.Unquote(q)
	if err != nil {
		return q
	}
	return result
}
