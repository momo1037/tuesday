package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 计算下一个周二的日期
	nextTuesday := nextWeekday(now, time.Tuesday)

	// 设置系统日期（保持时间不变）
	err := setSystemDate(nextTuesday)
	if err != nil {
		fmt.Printf("设置系统日期失败: %v\n", err)
		return
	}

	fmt.Printf("系统日期已成功修改为: %v\n", nextTuesday.Format("2006-01-02"))
}

// nextWeekday 计算给定日期之后的下一个指定星期几
func nextWeekday(t time.Time, weekday time.Weekday) time.Time {
	// 计算距离下一个指定星期几的天数
	daysUntil := (weekday - t.Weekday() + 7) % 7
	if daysUntil == 0 {
		daysUntil = 7 // 如果是当天，则跳到下周
	}
	return t.AddDate(0, 0, int(daysUntil))
}

// setSystemDate 设置Windows系统日期（保持时间不变）
func setSystemDate(newDate time.Time) error {
	// 转换为UTC时间以避免时区问题
	utcDate := newDate.UTC()
	// 准备系统时间结构
	systemTime := syscall.Systemtime{
		Year:  uint16(utcDate.Year()),
		Month: uint16(utcDate.Month()),
		Day:   uint16(utcDate.Day()),
		// 保持当前时间不变
		Hour:   uint16(utcDate.Hour()),
		Minute: uint16(utcDate.Minute()),
		Second: uint16(utcDate.Second()),
	}

	// 调用Windows API设置系统时间
	dll := syscall.NewLazyDLL("kernel32.dll")
	proc := dll.NewProc("SetSystemTime")
	r, _, err := proc.Call(uintptr(unsafe.Pointer(&systemTime)))
	if r == 0 {
		return fmt.Errorf("SetSystemTime调用失败: %v", err)
	}
	return nil
}
