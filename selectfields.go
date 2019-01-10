package dbhelper

import (
	"fmt"
	"reflect"
	"strings"
)

func SelectFields(e interface{}, prefix string) string {

	elements := reflect.ValueOf(e).Elem()
	var fields []string
	for i := 0; i < elements.NumField(); i++ {
		tag := elements.Type().Field(i).Tag
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

	return strings.Join(fields, ",")
}
