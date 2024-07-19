package util

import "strconv"

func StringToInt(s string) int {
	// 将字符转换为整数
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return atoi
}
