package index

import "github.com/bodagovsky/logs_out/tools"

/* Global index stores the data layout in memory for faster log files locating. It it unique per client*/

const MAXCHILDREN int = 6

type GlobalIndex struct {
	root *globalnode
}

type globalnode struct {
	value    int64
	pointer  int
	balanced bool
	children [MAXCHILDREN]*globalnode
	depth    int
}

func NewGlobalIndex() *GlobalIndex {
	return &GlobalIndex{
		root: &globalnode{
			children: [MAXCHILDREN]*globalnode{{}},
			depth:    1,
		},
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
	recursivelyInsert(newNode, gi.root, gi.root.depth+1 /* because we need a room to grow */)
}

func recursivelyInsert(node *globalnode, root *globalnode, leftDepth int) bool {
	if root.pointer == 0 || !root.children[0].balanced {
		if root.pointer == 0 {
			root.value = node.value
		}
		root.children[root.pointer] = node
		root.pointer++
		if root.pointer == MAXCHILDREN {
			if root.depth+1 > leftDepth {
				root.balanced = true
				return true
			}
			// otherwise - rebalance!
			rebalance(root)
		}
		return false
	}
	if root.children[root.pointer] == nil {
		root.children[root.pointer] = &globalnode{
			value: node.value,
			depth: 1,
		}
	}

	if recursivelyInsert(node, root.children[root.pointer], root.children[root.pointer-1].depth) {
		root.pointer++
		if root.pointer == MAXCHILDREN {
			if root.depth+1 > leftDepth {
				root.balanced = true
				return true
			}
			rebalance(root)
		}
	}

	return false
}

func rebalance(root *globalnode) {
	newNode := &globalnode{value: root.value, children: root.children, pointer: MAXCHILDREN, depth: root.children[0].depth + 1, balanced: true}
	root.children = [MAXCHILDREN]*globalnode{}
	root.children[0] = newNode
	root.pointer = 1
	root.depth++
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
