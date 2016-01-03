package helpers

import (
	"strconv"
)

func Str2Int(str string) int64 {
	integer, _ := strconv.ParseInt(str, 10, 64)
	return integer
}

/*func Interface2Int(v interface{}) (int64, error) {
	i, err :=
	return i, err
}*/
