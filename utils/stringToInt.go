package utils

import "strconv"

func StringToInt32(num string) int32  {
	value, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return int32(value)
}

func StringToInt64(num string) int64  {
	value, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return int64(value)
}
