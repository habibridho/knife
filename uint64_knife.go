package knife

type Uint64Knife struct {
	content []uint64
}

func NewUint64Knife(content []uint64) *Uint64Knife {
	return &Uint64Knife{content}
}

func (k *Uint64Knife) Filter(criteria func(item interface{}) bool) Knife {
	ret := []uint64{}
	for _, item := range k.content {
		if !criteria(item) {
			ret = append(ret, item)
		}
	}

	k.content = ret
	return k
}

func (k *Uint64Knife) Find(criteria func(item interface{}) bool) Knife {
	ret := []uint64{}
	for _, item := range k.content {
		if criteria(item) {
			ret = append(ret, item)
		}
	}

	k.content = ret
	return k
}

func (k *Uint64Knife) GetContent() []uint64 {
	return k.content
}
