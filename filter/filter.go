package filter

import (
	"errors"
	"reflect"
)

type Slicer interface {
	Filter(filter interface{}) Slicer
	Serve() []int
}

type slice struct {
	content interface{}
}

func NewSlicer(content interface{}) (Slicer, error) {
	if reflect.TypeOf(content).Kind() != reflect.Slice {
		return nil, errors.New("content is not a slice")
	}

	return &slice{content}, nil
}

func (s *slice) Filter(filter interface{}) Slicer {
	newContent := []int{}
	filterValue := reflect.ValueOf(filter)
	contentValue := reflect.ValueOf(s.content)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); content.Interface() != filterValue.Interface() {
			newContent = append(newContent, content.Interface().(int))
		}
	}

	s.content = newContent
	return s
}

func (s *slice) Serve() []int {
	return s.content.([]int)
}
