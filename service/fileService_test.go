package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestServiceImpl(t *testing.T) {
	f := &myFileSer{}
	tp := reflect.TypeOf(f)
	fmt.Println("--------------------")
	t.Logf(reflect.TypeOf((*FileService)(nil)).Elem().String())
	if tp.Implements(reflect.TypeOf((*FileService)(nil)).Elem()) {
		fmt.Println("Support !!!")
	}
}
