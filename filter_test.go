package knife

import "testing"

func TestFilterBool(t *testing.T) {
	slice := []bool{true, true, false, false, true}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(false).ServeBool()
	expectedResult := []bool{true, true, true}
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

func TestFilterInt8(t *testing.T) {
	slice := []int8{1, 2, 3, 4, 5, 5, 6}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(int8(5)).ServeInt8()
	expectedResult := []int8{1, 2, 3, 4, 6}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestFilterInt16(t *testing.T) {
	slice := []int16{1, 2, 3, 4, 5, 5, 6}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(int16(5)).ServeInt16()
	expectedResult := []int16{1, 2, 3, 4, 6}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestFilterint32(t *testing.T) {
	slice := []int32{1, 2, 3, 4, 5, 5, 6}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(int32(5)).ServeInt32()
	expectedResult := []int32{1, 2, 3, 4, 6}
	for i := range expectedResult {
		if result[i] != expectedResult[i] {
			t.Errorf("Not same, expected: %v, got: %v", expectedResult, result)
		}
	}
}

func TestFilterInt64(t *testing.T) {
	slice := []int64{1, 2, 3, 4, 5, 5, 6}
	slicer, _ := NewSlicer(slice)
	result := slicer.Filter(int64(5)).ServeInt64()
	expectedResult := []int64{1, 2, 3, 4, 6}
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

func TestSlice_FilterWithFunc(t *testing.T) {
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
	result := slicer.FilterWithFunc(func(currentValue interface{}) bool {
		return currentValue.(foo).key == "1"
	}).ServeInterface().([]foo)
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
