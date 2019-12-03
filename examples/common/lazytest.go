package common

import (
	"fmt"
	"os"
	"reflect"
)

func AssertNoError(err error) {
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
}

// Copied from Testify under MIT license.
// containsKind checks if a specified kind in the slice of kinds.
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

// Copied from Testify under MIT license.
func AssertNotNil(o interface{}) {
	if o == nil {
		fmt.Println("Object was nil!")
		os.Exit(255)
	}

	v := reflect.ValueOf(o)
	kind := v.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && v.IsNil() {
		fmt.Println("Object was nil!")
		os.Exit(255)
	}
}

func AssertNil(o interface{}) {
	if o == nil {
		return
	}

	v := reflect.ValueOf(o)
	kind := v.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && v.IsNil() {
		return
	}

	fmt.Println("Object nil!")
	os.Exit(255)
}

func AssertStringEqualsValue(a string, b string) {
	if a != b {
		fmt.Println("Strings weren't equal!")
		os.Exit(255)
	}
}

func AssertStringNotEqualsValue(a string, b string) {
	if a == b {
		fmt.Println("Strings weren't supposed to be equal!")
		os.Exit(255)
	}
}
