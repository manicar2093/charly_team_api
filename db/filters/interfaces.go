package filters

import "context"

type Filterable interface {
	// GetUserFilter looks up if the requested filter exists. If exists
	// Run method will be
	GetFilter(filterName string) error
	Run() (interface{}, error)
	SetContext(ctx context.Context)
	SetValues(values interface{})
}
