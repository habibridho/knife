package knife

type Server interface {
	ServeBool() []bool
	ServeString() []string
	ServeInt() []int
	ServeInt8() []int8
	ServeInt16() []int16
	ServeInt32() []int32
	ServeInt64() []int64
	ServeUint() []uint
	ServeUint8() []uint8
	ServeUint16() []uint16
	ServeUint32() []uint32
	ServeUint64() []uint64
	ServeByte() []byte
	ServeRune() []rune
	ServeFloat32() []float32
	ServeFloat64() []float64
	ServeComplex64() []complex64
	ServeComplex128() []complex128
	ServeInterface() interface{}
}

func (s *slice) ServeBool() []bool {
	return s.content.([]bool)
}

func (s *slice) ServeString() []string {
	return s.content.([]string)
}

func (s *slice) ServeInt() []int {
	return s.content.([]int)
}

func (s *slice) ServeInt8() []int8 {
	return s.content.([]int8)
}

func (s *slice) ServeInt16() []int16 {
	return s.content.([]int16)
}

func (s *slice) ServeInt32() []int32 {
	return s.content.([]int32)
}

func (s *slice) ServeInt64() []int64 {
	return s.content.([]int64)
}

func (s *slice) ServeUint() []uint {
	return s.content.([]uint)
}

func (s *slice) ServeUint8() []uint8 {
	return s.content.([]uint8)
}

func (s *slice) ServeUint16() []uint16 {
	return s.content.([]uint16)
}

func (s *slice) ServeUint32() []uint32 {
	return s.content.([]uint32)
}

func (s *slice) ServeUint64() []uint64 {
	return s.content.([]uint64)
}

func (s *slice) ServeByte() []byte {
	return s.content.([]byte)
}
func (s *slice) ServeRune() []rune {
	return s.content.([]rune)
}
func (s *slice) ServeFloat32() []float32 {
	return s.content.([]float32)
}
func (s *slice) ServeFloat64() []float64 {
	return s.content.([]float64)
}
func (s *slice) ServeComplex64() []complex64 {
	return s.content.([]complex64)
}
func (s *slice) ServeComplex128() []complex128 {
	return s.content.([]complex128)
}

func (s *slice) ServeInterface() interface{} {
	return s.content
}
