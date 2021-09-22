package injector

import (
	"reflect"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetNewObj(t *testing.T) {
	type test struct {
		t1 int
	}
	test1 := test{t1: 1}
	tc := newTransientContainer()

	test2 := tc.getNewObj(reflect.ValueOf(&test1)).Interface().(*test)
	test2.t1 = 2

	assert.NotEqual(t, test1, test2)
	assert.Equal(t, test1.t1, 1)
	assert.Equal(t, test2.t1, 2)
}
