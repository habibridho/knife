package knife

import "testing"

func TestUint64Knife(t *testing.T) {
	knife := NewUint64Knife([]uint64{1, 2, 3, 4, 5, 6, 7, 7, 8, 12, 3, 2, 1, 2})
	knife.Filter(func(item interface{}) bool {
		return item.(uint64) == 2
	})
	knife.Find(func(item interface{}) bool {
		return item.(uint64)%2 == 0
	})
	result := knife.GetContent()

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

func TestStringKnife(t *testing.T) {
	knife := NewStringKnife([]string{"1", "2", "3", "4", "5", "6", "7", "7", "8", "12", "3", "2", "1", "2"})
	knife.Filter(func(item interface{}) bool {
		return item.(string) == "2"
	})
	knife.Find(func(item interface{}) bool {
		return item.(string) == "12"
	})
	result := knife.GetContent()

	expectedResult := []string{"12"}
	if len(result) != len(expectedResult) {
		t.Fatalf("Result length not expected")
	}
	for i := range result {
		if result[i] != expectedResult[i] {
			t.Fatalf("Unexpected result")
		}
	}
}
