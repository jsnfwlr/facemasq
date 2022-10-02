package utils

import (
	"strconv"
	"strings"
)

func CSVtoInt64s(csv string) (nums []int64, err error) {
	var num int
	strs := strings.Split(csv, ",")
	for _, str := range strs {
		num, err = strconv.Atoi(str)
		nums = append(nums, int64(num))
	}
	return
}
