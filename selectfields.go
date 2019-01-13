package dbhelper

import (
	"fmt"
	"reflect"
	"strings"
)

func SelectFields(e interface{}, prefix string) string {

	elements := reflect.ValueOf(e).Elem()
	fields := transfromFields(elements, prefix)

	return strings.Join(fields, ",")
}

func transfromFields(f reflect.Value, prefix string) []string {

	var fields []string
	for i := 0; i < f.NumField(); i++ {
		tag := f.Type().Field(i).Tag
		v, ok := tag.Lookup("db")

		if ok != false {
			if v != "" && v != "-" {
				if prefix != "" {
					v = fmt.Sprintf("%v.%v", prefix, v)
				}
				fields = append(fields, v)
			}
		}
	}

	return fields
}
