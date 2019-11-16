package filter

import "testing"

func TestFilterInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 5, 6}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(5).ServeInt()
	expectedResult := []int{1, 2, 3, 4, 6}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestFilterString(t *testing.T) {
	slice := []string{"abc", "cde", "a", "a", "a", "foo", "bar"}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter("a").ServeString()
	expectedResult := []string{"abc", "cde", "foo", "bar"}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestFilterStruct(t *testing.T) {
	type foo struct {
		key string
		val int
	}
	slice := []foo{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
	}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(foo{"1", 1}).ServeInterface().([]foo)
	expectedResult := []foo{
		{"2", 2},
		{"3", 3},
		{"4", 4},
	}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestWrongInput(t *testing.T) {
	slice := "some wrong slice"
	_, err := NewSlicer(slice)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
}
