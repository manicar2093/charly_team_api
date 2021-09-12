package filters

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilterTests struct {
	suite.Suite
	filterParams      *FilterParameters
	filterData1       FilterRegistrationData
	filterData2       FilterRegistrationData
	filterData3       FilterRegistrationData
	filtersToRegister []FilterRegistrationData
}

func (c *FilterTests) SetupTest() {
	c.filterParams = &FilterParameters{}
	c.filterData1 = FilterRegistrationData{Name: "filter1", Func: func(filterParameters *FilterParameters) (interface{}, error) { return "runned", nil }}
	c.filterData2 = FilterRegistrationData{Name: "filter2", Func: func(filterParameters *FilterParameters) (interface{}, error) { return "runned", nil }}
	c.filterData3 = FilterRegistrationData{Name: "filter3", Func: func(filterParameters *FilterParameters) (interface{}, error) { return "runned", nil }}
	c.filtersToRegister = []FilterRegistrationData{c.filterData1, c.filterData2, c.filterData3}
}

func (c *FilterTests) TearDownTest() {

}

func (c *FilterTests) TestFilter_GetFilter() {
	filter := NewFilter(c.filterParams, c.filtersToRegister...)
	c.Nil(filter.GetFilter("filter1"), "should be found")
}

func (c *FilterTests) TestFilter_FilterNotFound() {
	filter := NewFilter(c.filterParams, c.filtersToRegister...)
	c.NotNil(filter.GetFilter("not_exists"), "should not be found")
}

func (c *FilterTests) TestFilter_Run() {
	filter := NewFilter(c.filterParams, c.filtersToRegister...)
	c.Nil(filter.GetFilter("filter1"), "should be found")
	got, err := filter.Run()
	c.Nil(err, "should not get an error")
	c.Equal("runned", got.(string), "bad filter response")
}

func (c *FilterTests) TestFilter_SetCtx() {
	filter := NewFilter(c.filterParams, c.filtersToRegister...)
	c.Nil(filter.GetFilter("filter1"), "should be found")
	ctx := context.Background()
	filter.SetContext(ctx)
	filterImpl := filter.(*Filter)
	c.Equal(filterImpl.filterParams.Ctx, ctx)
}

func (c *FilterTests) TestFilter_SetValues() {
	filter := NewFilter(c.filterParams, c.filtersToRegister...)
	c.Nil(filter.GetFilter("filter1"), "should be found")
	newValues := "new-values"
	filter.SetValues(newValues)
	filterImpl := filter.(*Filter)
	c.Equal(filterImpl.filterParams.Values, newValues)
}

func TestFilters(t *testing.T) {
	suite.Run(t, new(FilterTests))
}
