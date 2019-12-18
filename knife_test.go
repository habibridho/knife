package knife

import "testing"

func TestSlicer(t *testing.T) {
	slicer := NewSlicer([]uint64{1, 2, 3, 4, 5, 6, 7, 7, 8, 12, 3, 2, 1, 2})
	slicer.Filter(func(item uint64) bool {
		return item == 2
	})
	slicer.Find(func(item uint64) bool {
		return item%2 == 0
	})
	result := slicer.GetContent()

	expectedResult := []uint64{4, 6, 8, 12}
	if len(result) != len(expectedResult) {
		t.Fatalf("Result length not expected")
	}
	for i := range result {
		if result[i] != expectedResult[i] {
			t.Fatalf("Unexpected result")
		}
	}
}
