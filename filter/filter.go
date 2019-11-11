package filter

type Slicer interface {
	Filter(filter int) Slicer
	Serve() []int
}

type slice struct {
	content []int
}

func NewSlicer(content []int) Slicer {
	return &slice{content}
}

func (s *slice) Filter(filter int) Slicer {
	newContent := []int{}
	for _, num := range s.content {
		if num != filter {
			newContent = append(newContent, num)
		}
	}

	s.content = newContent
	return s
}

func (s *slice) Serve() []int {
	return s.content
}
