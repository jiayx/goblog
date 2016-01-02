package helpers

import (
	"strconv"
)

func Str2Int(str string) (int64, error) {
	integer, err := strconv.ParseInt(str, 10, 64)
	return integer, err
}
