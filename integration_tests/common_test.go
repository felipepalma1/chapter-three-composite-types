package integration_tests

import (
	"github.com/felipepalma1/chapter-three-composite-types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleArray(t *testing.T) {
	var targetResult [3]int
	functionResult, _ := chapter_three_composite_types.SimpleArray()

	assert.Equal(t, functionResult, targetResult)
}

func TestSimpleInitializedArray(t *testing.T) {
	var targetResult = [3]int{10, 20, 30}
	functionResult := chapter_three_composite_types.SimpleInitializedArray()

	assert.Equal(t, functionResult, targetResult)
}

func TestNewSpecialArray(t *testing.T) {
	var expectedValues = [12]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1}
	x := chapter_three_composite_types.NewSpecialArray()
	assert.Equal(t, expectedValues, x)
}

func TestLoadingASlice(t *testing.T) {
	x := chapter_three_composite_types.LoadingASlice()
	assert.NotNil(t, x)
}

func TestMakingSlice(t *testing.T) {
	x := chapter_three_composite_types.MakingSlice()
	assert.NotNil(t, x)
}
