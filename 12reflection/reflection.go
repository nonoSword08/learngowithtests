package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(v interface{}) reflect.Value {
	val := reflect.ValueOf(v)

	// 如果是val是指针，需要先把指针指向的值取出来
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
