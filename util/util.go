package util

import (
	"bytes"
	"encoding/json"
	// "fmt"
	// "reflect"
	"strconv"
	"strings"
	"time"
)

const Separator = ","

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

func SkipPeriodTime(start time.Time, d time.Duration) time.Time {
	return skipPeriodTime3(time.Now(), start, d)
}

func skipPeriodTime3(now, start time.Time, d time.Duration) time.Time {
	end := start
	if diff := now.Sub(start); diff > 0 && d > 0 {
		end = start.Add(time.Duration((diff + d - 1) / d * d))
	}
	return end
}

// Monday,Thursday...
func GetFirstWeekday(t time.Time) time.Time {
	weekDays := ((int)(t.Weekday()) + 6) % 7
	firstDay := t.Add(-time.Duration(weekDays) * 24 * time.Hour)
	firstDay, _ = ParseTime(firstDay.Format("2006-01-02"))
	return firstDay
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

// compare a,b json string
// TODO  ignore struct or map field order
func DeepEqual(a, b interface{}) bool {
	b1, _ := json.Marshal(a)
	b2, _ := json.Marshal(b)
	return bytes.Compare(b1, b2) == 0
}
