package index

import "github.com/bodagovsky/logs_out/tools"

/* Global index stores the data layout in memory for faster log files locating. It it unique per client*/

type GlobalIndex struct {
	root *globalnode
}

type globalnode struct {
	value    int64
	children []*globalnode
}

func (gi *GlobalIndex) Insert(timestamp int64) {
	if gi.root.value > timestamp {
		return
	}
	newNode := &globalnode{value: timestamp}
	recursivelyInsert(newNode, gi.root.children)
}

func recursivelyInsert(node *globalnode, children []*globalnode) {}

// CompareGlobalNode compares two globalnode pointers and returns a tools.Comparison value.
func CompareGlobalNode(a, b *globalnode) tools.Comparison {
	switch {
	case a.value > b.value:
		return tools.GREATER
	case a.value < b.value:
		return tools.LESS
	default:
		return tools.EQ
	}
}
