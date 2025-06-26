package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalIndex(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1_750_921_338

	for range 36 {
		globalIndex.Insert(ts)
		ts++
	}

	assert.Equal(t, int64(1_750_921_338), globalIndex.root.value)

	assert.Equal(t, int64(1_750_921_338), globalIndex.root.children[0].children[0].value)
	assert.Equal(t, int64(1_750_921_339), globalIndex.root.children[0].children[1].value)
	assert.Equal(t, int64(1_750_921_340), globalIndex.root.children[0].children[2].value)
	assert.Equal(t, int64(1_750_921_341), globalIndex.root.children[0].children[3].value)
	assert.Equal(t, int64(1_750_921_342), globalIndex.root.children[0].children[4].value)
	assert.Equal(t, int64(1_750_921_343), globalIndex.root.children[0].children[5].value)

	assert.Equal(t, int64(1_750_921_344), globalIndex.root.children[1].value)
	assert.Equal(t, int64(1_750_921_344), globalIndex.root.children[1].children[0].value)

	assert.Equal(t, 2, globalIndex.overallDepth)
}
