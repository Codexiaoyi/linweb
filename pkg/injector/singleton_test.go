package injector

import (
	"reflect"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSingletonObject(t *testing.T) {
	type test struct {
		t1 int
	}
	test1 := test{t1: 1}
	sc := newSingletonContainer()
	sc.setObject(&test1)
	getV := sc.getObject(reflect.TypeOf(&test1))
	assert.Equal(t, getV, reflect.ValueOf(&test1))
	getV1 := sc.getObject(reflect.TypeOf(&test1))
	assert.Equal(t, getV1, reflect.ValueOf(&test1))
}
