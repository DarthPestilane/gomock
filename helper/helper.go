package helper

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func Dump(v ...interface{}) {
	fmt.Printf("%+v\n", v)
}

func RandomSlice(array interface{}, valPtr interface{}) {
	arrVal := reflect.Indirect(reflect.ValueOf(array))
	if arrVal.Kind() != reflect.Slice {
		panic(fmt.Errorf("first param expected a slice not %s", arrVal.Kind().String()))
	}
	val := reflect.ValueOf(valPtr)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("second param expected a ptr not %s", val.Kind().String()))
	}
	if arrVal.Len() == 0 {
		return
	}
	idx := Random(arrVal.Len())
	val.Elem().Set(arrVal.Index(idx))
}

func Random(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
