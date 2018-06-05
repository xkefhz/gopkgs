package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	Separator = ","
)

func ParseTime(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	match := "2006-01-02 15:04:05"
	if form := "2006-01-02"; len(form) == len(s) {
		match = form
	}
	if form := "2006-1-2"; len(form) == len(s) {
		match = form
	}
	return time.ParseInLocation(match, s, loc)
}

func SkipPeriodTime(startTime time.Time, d time.Duration) time.Time {
	endTime := startTime
	if diff := time.Now().Sub(startTime); diff > 0 && d > 0 {
		endTime = startTime.Add(time.Duration((diff + d - 1) / d * d))
	}
	return endTime
}

func ParseStrings(s string) []string {
	s = strings.Replace(s, ";", ",", -1)
	s = strings.Replace(s, "-", ",", -1)
	s = strings.Replace(s, "~", ",", -1)
	s = strings.Replace(s, "/", ",", -1)
	s = strings.Replace(s, "\\", ",", -1)
	return strings.Split(s, ",")
}

func ParseIntSlice(s string) []int64 {
	chips := make([]int64, 0, 8)
	for _, v := range ParseStrings(s) {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			chips = append(chips, n)
		}
	}
	return chips
}

func ContainsBySeparator(s, sep, obj string) bool {
	s = strings.Replace(s, ";", ",", -1)
	parts := strings.Split(s, sep)
	for _, part := range parts {
		if part == obj {
			return true
		}
	}
	return false
}

func FormatMoney(money int64) string {
	var s, prefix string
	if money == 0 {
		s = "0"
	} else if money < 0 {
		prefix = "-"
		money = -money
	}
	for money > 0 {
		mod := money % 1000
		money /= 1000
		if money > 0 {
			s = "," + fmt.Sprintf("%03d", mod) + s
		} else {
			s = fmt.Sprintf("%d", mod) + s
		}
	}
	s = prefix + s
	return s
}

func InArray(array interface{}, some interface{}) int {
	counter := 0
	someValues := reflect.ValueOf(some)
	arrayValues := reflect.ValueOf(array)
	for i := 0; i < arrayValues.Len(); i++ {
		if someValues.Kind() == reflect.Slice {
			for k := 0; k < someValues.Len(); k++ {
				if reflect.DeepEqual(arrayValues.Index(i).Interface(), someValues.Index(k).Interface()) == true {
					counter++
				}
			}
		}
		if reflect.DeepEqual(arrayValues.Index(i).Interface(), some) {
			counter++
		}
	}
	return counter
}

// 素数
func IsPrimeNumber(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}