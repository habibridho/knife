package knife

type slice struct {
	content []uint64
}

func NewSlicer(content []uint64) *slice {
	return &slice{content}
}

func (s *slice) Filter(f func(item uint64) bool) *slice {
	ret := []uint64{}
	for _, item := range s.content {
		if !f(item) {
			ret = append(ret, item)
		}
	}

	s.content = ret
	return s
}

func (s *slice) Find(f func(item uint64) bool) *slice {
	ret := []uint64{}
	for _, item := range s.content {
		if f(item) {
			ret = append(ret, item)
		}
	}

	s.content = ret
	return s
}

func (s *slice) GetContent() []uint64 {
	return s.content
}
