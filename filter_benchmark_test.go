package knife

import "testing"

func BenchmarkSlice_Filter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []string{"abc", "cde", "a", "a", "a", "foo", "bar"}
		slicer, _ := NewSlicer(slice)
		slicer.Filter("a").ServeString()
	}
}

func BenchmarkRawFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []string{"abc", "cde", "a", "a", "a", "foo", "bar"}
		rawfilter(slice, "a")
	}
}

func rawfilter(slice []string, filter string) []string {
	result := []string{}
	for i := range slice {
		if slice[i] != filter {
			result = append(result, slice[i])
		}
	}

	return result
}
