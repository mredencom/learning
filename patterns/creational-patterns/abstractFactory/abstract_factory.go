package abstractFactory

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	MemoryStorage
	TempStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case DiskStorage:
		return newDiskStorage()
	case MemoryStorage:
		return newMemoryStorage()
	default:
		return newTempStorage()
	}
}

type Disk struct{}

func newDiskStorage() *Disk {
	return &Disk{}
}

func (s *Disk) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Memory struct{}

func newMemoryStorage() *Memory {
	return &Memory{}
}

func (s *Memory) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Temp struct{}

func newTempStorage() *Temp {
	return &Temp{}
}

func (s *Temp) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}
