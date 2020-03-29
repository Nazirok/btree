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

// Search search key in node
func (bt *Btree) Search(key int64) {
	bt.search(bt.root, key)
}

func (bt *Btree) search(node *bNode, key int64) {
	var i int64
	for i <= node.kcount && key > node.keys[i] {
		i++
	}
	if i <= node.kcount && key == node.keys[i] {
		return
	} else if node.leaf {
		return
	} else {
		// TODO: diskRead(node.childs[i])
		// return bt.search()
	}
}
