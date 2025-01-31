package testinstance

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Set sets the value of the field with the given name on the target struct.
// It panics if the field is not found or if the target is not a struct.
// It is used to set the values of the unexported fields of a struct in the test cases.
func Set[T any](target T, fields map[string]any) T {
	for fieldName, value := range fields {
		setFieldValue(target, fieldName, value)
	}
	return target
}

func setFieldValue(target any, fieldName string, value any) {
	rv := reflect.ValueOf(target)
	for rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}

	if !rv.CanAddr() {
		panic("target must be addressable")
	}

	if rv.Kind() != reflect.Struct {
		panic(fmt.Sprintf(
			"unable to set the '%s' field value of the type %T, target must be a struct",
			fieldName,
			target,
		))
	}

	rf := rv.FieldByName(fieldName)
	reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}
