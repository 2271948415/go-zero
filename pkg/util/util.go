package util

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func RandomNumeric(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		panic("{ size : " + strconv.Itoa(size) + " } must be more than 0 ")
	}
	value := ""
	for index := 0; index < size; index++ {
		value += strconv.Itoa(r.Intn(10))
	}

	return value
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// IsEmpty 判断是否为空
func IsEmpty(data interface{}) bool {
	if data == nil {
		return true
	}
	dataRef := reflect.ValueOf(data)
	for dataRef.Kind() == reflect.Ptr {
		dataRef = dataRef.Elem()
	}
	switch dataRef.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return dataRef.Len() == 0
	case reflect.String:
		return strings.Trim(data.(string), " ") == ""
	}
	return false
}
