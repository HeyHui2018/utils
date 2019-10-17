package utils

import (
	"time"
)

/*
通过输入的星期获取当天日期
本函数以星期天为一周的起点
*/

func GetDateByWeekday(weekday time.Weekday) time.Time {
	today := time.Now()
	return today.Add(-time.Duration(today.Weekday() - weekday) * 24 * time.Hour)
}
