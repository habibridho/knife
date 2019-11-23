package knife

import (
	"errors"
	"reflect"
)

type Slicer interface {
	Filterer
	Server
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
