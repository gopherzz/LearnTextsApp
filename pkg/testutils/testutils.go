package testutils

import (
	"log"
	"reflect"
)

// Simple AssertEqual crutch
func AssertEqual(excepted, get interface{}) {
	if !reflect.DeepEqual(excepted, get) {
		log.Fatalf("Error - get: %v, excepted: %v", get, excepted)
	}
}
