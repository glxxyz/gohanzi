package repo_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Based on https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func assert(t testing.TB, condition bool, message string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+message+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		t.FailNow()
	}
}

// assertOK fails the test if an err is not nil.
func assertOK(t testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		t.FailNow()
	}
}

// assert_equals fails the test if expected is not equal to actual.
func assertEquals(t testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texpected: %#v\n\n\tactual: %#v\033[39m\n\n", filepath.Base(file), line, expected, actual)
		t.FailNow()
	}
}
