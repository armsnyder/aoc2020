package common

import "reflect"

func AssertEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic("assert failed")
	}
}
