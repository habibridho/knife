package knife

type StringKnife struct {
	content []string
}

func NewStringKnife(content []string) *StringKnife {
	return &StringKnife{content}
}

func (k *StringKnife) Filter(criteria func(item interface{}) bool) Knife {
	ret := []string{}
	for _, item := range k.content {
		if !criteria(item) {
			ret = append(ret, item)
		}
	}

	k.content = ret
	return k
}

func (k *StringKnife) Find(criteria func(item interface{}) bool) Knife {
	ret := []string{}
	for _, item := range k.content {
		if criteria(item) {
			ret = append(ret, item)
		}
	}

	k.content = ret
	return k
}

func (k *StringKnife) GetContent() []string {
	return k.content
}
