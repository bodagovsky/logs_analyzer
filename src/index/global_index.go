package index

import "github.com/bodagovsky/logs_out/tools"

/* Global index stores the data layout in memory for faster log files locating. It it unique per client*/

const MAXCHILDREN int = 6

type GlobalIndex struct {
	root         *globalnode
	overallDepth int
}

type globalnode struct {
	value      int64
	pointer    int
	rebalanced bool
	children   [MAXCHILDREN]*globalnode
}

func NewGlobalIndex() *GlobalIndex {
	return &GlobalIndex{
		root: &globalnode{},
	}
}

func (gi *GlobalIndex) Insert(timestamp int64) {
	if gi.root.value == 0 {
		gi.root.value = timestamp
	}
	if gi.root.value > timestamp {
		return
	}
	newNode := &globalnode{value: timestamp}
	if recursivelyInsert(newNode, gi.root, 1, gi.overallDepth) {
		rebalance(gi.root)
		gi.overallDepth++
	}
}

func recursivelyInsert(node *globalnode, root *globalnode, currDepth int, overallDepth int) bool {
	if root.pointer == 0 || !root.children[root.pointer-1].rebalanced {
		if root.pointer == 0 {
			root.value = node.value
		}
		root.children[root.pointer] = node
		root.pointer++
		if root.pointer == MAXCHILDREN {
			if currDepth == overallDepth {
				return true
			}
			// otherwise - rebalance!
			rebalance(root)
		}
		return false
	}
	if root.children[root.pointer] == nil {
		root.children[root.pointer] = &globalnode{}
	}
	if recursivelyInsert(node, root.children[root.pointer], currDepth+1, overallDepth) {
		root.pointer++
		if root.pointer == MAXCHILDREN {
			rebalance(root)
		}
	}

	return false
}

func rebalance(root *globalnode) {
	newNode := &globalnode{value: root.value, children: root.children, rebalanced: true}
	root.children = [MAXCHILDREN]*globalnode{}
	root.children[0] = newNode
	root.pointer = 1
}

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
