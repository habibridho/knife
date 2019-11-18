package filter

import (
	"errors"
	"reflect"
)

type Slicer interface {
	Filter(filter interface{}) Slicer
	FilterWithFunc(f func(currentValue interface{}) bool) Slicer
	ServeInt() []int
	ServeString() []string
	ServeInterface() interface{}
}

type slice struct {
	content     interface{}
	contentType reflect.Type
}

func NewSlicer(content interface{}) (Slicer, error) {
	contentType := reflect.TypeOf(content)
	if contentType.Kind() != reflect.Slice {
		return nil, errors.New("content is not a slice")
	}

	return &slice{content, contentType}, nil
}

func (s *slice) Filter(filter interface{}) Slicer {
	filterValue := reflect.ValueOf(filter)
	contentValue := reflect.ValueOf(s.content)
	newContent := reflect.MakeSlice(s.contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); content.Interface() != filterValue.Interface() {
			newContent = reflect.Append(newContent, content)
		}
	}

	s.content = newContent.Interface()
	return s
}

func (s *slice) FilterWithFunc(f func(currentValue interface{}) bool) Slicer {
	contentValue := reflect.ValueOf(s.content)
	newContent := reflect.MakeSlice(s.contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); !f(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}

	s.content = newContent.Interface()
	return s
}

func (s *slice) ServeInt() []int {
	return s.content.([]int)
}

func (s *slice) ServeString() []string {
	return s.content.([]string)
}

func (s *slice) ServeInterface() interface{} {
	return s.content
}
