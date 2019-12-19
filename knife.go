package knife

type Knife interface {
	// Filter out items that match criteria
	Filter(criteria func(item interface{}) bool) Knife

	// Find items that match criteria
	Find(criteria func(item interface{}) bool) Knife
}
