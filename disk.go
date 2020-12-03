package btree

import "os"

func newDiskService(dbPath string) (*diskService, error) {
	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &diskService{file: file}, nil
}

type diskService struct {
	file *os.File
}

//
func (ds *diskService) readNode(blockID uint64) (*node, error) {
	return nil, nil
}
