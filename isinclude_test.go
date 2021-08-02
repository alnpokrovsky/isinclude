package isinclude_test

import (
	"testing"

	"github.com/alnpokrovsky/isinclude"
	"github.com/stretchr/testify/assert"
)

func TestIsIncludeBasic(t *testing.T) {
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 5, 7, 9, 11}, []int{}))
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 5, 7, 9, 11}, []int{3, 5, 7}))
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 5, 7, 9, 11}, []int{4, 5, 7}))
}

func TestIsIncludeDuplicates(t *testing.T) {
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 3, 5, 7, 9, 11}, []int{3, 3, 5, 7}))
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 3, 3, 5, 7, 9, 11}, []int{3, 3, 5, 7}))
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 3, 3, 3, 3, 5, 7, 9, 11}, []int{3, 3, 5, 7}))
	assert.Equal(t, true, isinclude.IsInclude([]int{1, 2, 3, 3, 3, 5, 7, 9, 11}, []int{3, 5, 7}))
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 3, 3, 5, 5, 7, 9, 11}, []int{3, 5, 7}))
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 3, 3, 5, 5, 5, 5, 5, 5, 7, 7, 8, 9, 11}, []int{3, 5, 7}))
}

func TestIsIncludeSpaces(t *testing.T) {
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 5, 7, 9, 11}, []int{3, 5, 9}))
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 5, 5, 5, 7, 9, 11}, []int{5, 5, 9}))
	assert.Equal(t, false, isinclude.IsInclude([]int{1, 2, 3, 5, 7, 9, 11}, []int{3, 11}))
}

func TestIsIncludeReverce(t *testing.T) {
	assert.Equal(t, true, isinclude.IsInclude([]int{11, 10, 9, 5, 4, 3, 1}, []int{5, 4, 3}))
	assert.Equal(t, false, isinclude.IsInclude([]int{11, 10, 9, 5, 4, 3, 1}, []int{5, 4, 2}))
	assert.Equal(t, true, isinclude.IsInclude([]int{11, 10, 9, 9, 9, 9, 5, 4, 3, 1}, []int{9, 9, 5, 4, 3}))
	assert.Equal(t, true, isinclude.IsInclude([]int{11, 10, 9, 9, 5, 4, 3, 1}, []int{9, 9, 5, 4, 3}))
	assert.Equal(t, false, isinclude.IsInclude([]int{11, 10, 9, 5, 4, 3, 1}, []int{9, 9, 5, 4, 3}))
}
