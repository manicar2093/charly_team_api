package filters

type FilterService interface {
	// GetUserFilter looks up if the requested filter exists. If exists
	// Run method will be
	GetUserFilter(string) FilterRunable
}

type FilterRunable interface {
	Run(filterParameters *FilterParameters) (interface{}, error)
	IsFound() bool
}
