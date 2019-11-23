package knife

import (
	"reflect"
)

type Filterer interface {
	Filter(filter interface{}) Slicer
	FilterWithFunc(f func(currentValue interface{}) bool) Slicer
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
