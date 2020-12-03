package btree

const t = 5

// Btree B-tree struct
type Btree struct {
	diskService *diskService
	root        *node
}

// Search search key in node
func (bt *Btree) Search(key uint64) error {
	return bt.search(bt.root, key)
}

func (bt *Btree) search(node *node, key uint64) error {
	var i uint64
	for i <= node.keysCnt && key > node.keys[i].key {
		i++
	}
	if i <= node.keysCnt && key == node.keys[i].key {
		return nil
	} else if node.leaf {
		return nil
	} else {
		childNode, err := bt.diskService.readNode(node.childrenBlockIDs[i])
		if err != nil {
			return err
		}
		return bt.search(childNode, key)
	}
}
