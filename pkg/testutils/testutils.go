package testutils

import (
	"log"
	"reflect"
)

func Assert(excepted, get interface{}) {
	if !reflect.DeepEqual(excepted, get) {
		log.Fatalf("Error - get: %v, excepted: %v", get, excepted)
	}
}
