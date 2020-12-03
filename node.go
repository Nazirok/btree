package btree

// 2+8+90 = 103
const kvSize = 103

type (
	node struct {
		keysCnt          uint64
		keys             [2*t - 1]kv
		childrenBlockIDs [2 * t]uint64
		blockID          uint64
		leaf             bool
	}

	// total 124 bytes
	kv struct {
		valueLen uint16 // 2
		key      uint64 // 8
		value    string // 93
	}
)
