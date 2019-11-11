package filter

import "testing"

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 5, 6}
	slicer := NewSlicer(slice)
	result := slicer.Filter(5).Serve()
	expectedResult := []int{1, 2, 3, 4, 6}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same")
		}
	}
}
