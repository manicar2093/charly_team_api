package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testingUpdateRequestValidStruct struct {
	ID   int32
	UUID string
}

func (c testingUpdateRequestValidStruct) GetID() int32 {
	return c.ID
}

func (c testingUpdateRequestValidStruct) GetUUID() string {
	return c.UUID
}

func TestIsUpdateRequestValid_BothData(t *testing.T) {
	testData := testingUpdateRequestValidStruct{ID: 1, UUID: "uuid"}
	got := IsUpdateRequestValid(testData)
	assert.True(t, got, "Struct should be valid")
}

func TestIsUpdateRequestValid_OnlyID(t *testing.T) {
	testData := testingUpdateRequestValidStruct{ID: 1}
	got := IsUpdateRequestValid(testData)
	assert.True(t, got, "Struct should be valid")
}

func TestIsUpdateRequestValid_OnlyUUID(t *testing.T) {
	testData := testingUpdateRequestValidStruct{UUID: "uuid"}
	got := IsUpdateRequestValid(testData)
	assert.True(t, got, "Struct should be valid")
}

func TestIsUpdateRequestValid_NoData(t *testing.T) {
	testData := testingUpdateRequestValidStruct{}
	got := IsUpdateRequestValid(testData)
	assert.False(t, got, "Struct should be invalid")
}
