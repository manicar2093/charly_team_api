package filters

import (
	"testing"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestFilterRunner(t *testing.T) {

	params := FilterParameters{}
	mockedFuncItemReturned := "a returner value"
	mockedFunc := func(filterParameters *FilterParameters) (interface{}, error) { return mockedFuncItemReturned, nil }

	runner := FilterRunner{FilterName: "a filter", Filter: mockedFunc, Found: true}

	assert.True(t, runner.IsFound(), "function should be found")

	itemGot, errGot := runner.Run(&params)

	assert.Nil(t, errGot, "should not return an error")

	itemGotAsString, ok := itemGot.(string)
	assert.True(t, ok, "error parsing filter response")
	assert.Equal(t, mockedFuncItemReturned, itemGotAsString, "item response unexpected")

}

func TestFilterRunner_NotFound(t *testing.T) {

	params := FilterParameters{}

	runner := FilterRunner{FilterName: "a filter", Found: false}

	assert.False(t, runner.IsFound(), "function should be found")

	itemGot, errGot := runner.Run(&params)

	assert.Nil(t, itemGot, "should not return an item")

	_, isBadStatusError := errGot.(apperrors.BadStatusError)
	assert.True(t, isBadStatusError, "unexpected type of error")

}
