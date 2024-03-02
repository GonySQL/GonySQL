package utils

import "reflect"

func ToType(type_arr []string) []reflect.Type {
	typeResult := make([]reflect.Type, len(type_arr))
	for i, v := range type_arr {
		switch v {
		case "int":
			typeResult[i] = reflect.TypeOf(0)
		case "string":
			typeResult[i] = reflect.TypeOf("")
		}
	}
	return typeResult
}
