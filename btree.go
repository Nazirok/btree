package btree

const t = 5

// Btree B-tree struct
type Btree struct {
	root *bNode
}

type bNode struct {
	kcount int64
	scount int64
	keys   [2*t - 1]int64
	childs [2 * t]*bNode
	parent *bNode
	leaf   bool
}
