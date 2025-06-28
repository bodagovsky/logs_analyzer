package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalIndex_Insert_6(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1_750_921_338
	start := ts

	for range 5 {
		globalIndex.Insert(ts)
		ts++
	}
	assert.Equal(t, 1, globalIndex.root.depth)
	assert.Equal(t, int64(1_750_921_338), globalIndex.root.value)

	globalIndex.Insert(ts)

	assert.Equal(t, 2, globalIndex.root.depth)

	step := int64(1)
	for _, child := range globalIndex.root.children[0].children {
		assert.Equal(t, start, child.value)
		start += step
	}
}

func TestGlobalIndex_Insert_36(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1_750_921_338
	start := ts

	for range 35 {
		globalIndex.Insert(ts)
		ts++
	}

	assert.Equal(t, int64(1_750_921_338), globalIndex.root.value)
	assert.Equal(t, 2, globalIndex.root.depth)
	globalIndex.Insert(ts)
	assert.Equal(t, 3, globalIndex.root.depth)

	step := int64(6)
	children := globalIndex.root.children[0].children

	for range globalIndex.root.depth - 1 {
		for _, child := range children {
			assert.Equal(t, start, child.value)
			start += step
		}
		start -= int64(len(children) * int(step))
		step /= 6
		children = children[0].children
	}

}

func TestGlobalIndex_Insert_216(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1
	start := ts

	for range 215 {
		globalIndex.Insert(ts)
		ts++
	}

	assert.Equal(t, 3, globalIndex.root.depth)
	globalIndex.Insert(ts)
	assert.Equal(t, 4, globalIndex.root.depth)

	step := int64(36)
	children := globalIndex.root.children[0].children

	for range globalIndex.root.depth - 1 {
		for _, child := range children {
			assert.Equal(t, start, child.value)
			start += step
		}
		start -= int64(len(children) * int(step))
		step /= 6
		children = children[0].children
	}
}

func TestGlobalIndex_Insert_1296(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1
	start := ts

	for range 1295 {
		globalIndex.Insert(ts)
		ts++
	}
	assert.Equal(t, 4, globalIndex.root.depth)
	globalIndex.Insert(ts)
	assert.Equal(t, 5, globalIndex.root.depth)

	step := int64(216)
	children := globalIndex.root.children[0].children

	for range globalIndex.root.depth - 1 {
		for _, child := range children {
			assert.Equal(t, start, child.value)
			start += step
		}
		start -= int64(len(children) * int(step))
		step /= 6
		children = children[0].children
	}
}

func TestGlobalIndex_Insert_7776(t *testing.T) {
	globalIndex := NewGlobalIndex()

	var ts int64 = 1
	start := ts

	for range 7775 {
		globalIndex.Insert(ts)
		ts++
	}
	assert.Equal(t, 5, globalIndex.root.depth)
	globalIndex.Insert(ts)
	assert.Equal(t, 6, globalIndex.root.depth)

	step := int64(1296)
	children := globalIndex.root.children[0].children

	for range globalIndex.root.depth - 1 {
		for _, child := range children {
			assert.Equal(t, start, child.value)
			start += step
		}
		start -= int64(len(children) * int(step))
		step /= 6
		children = children[0].children
	}
}
