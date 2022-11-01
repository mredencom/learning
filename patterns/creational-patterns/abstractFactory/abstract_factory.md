# 工厂方法

工厂方法设计模式是允许创建对象而无需指定将要创建的对象的确切类型。

## 接口实现

接口展现存储数据，后台实现不同的方式：diskStorage , memoryStorage, tempStorage

### 接口定义

```go
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
```

### 接口实现


```go
type Disk struct{}

func newDiskStorage() *Disk {
	return &Disk{}
}

func (s *Disk) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}
```

```go
type Memory struct{}

func newMemoryStorage() *Memory {
	return &Memory{}
}

func (s *Memory) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}
```

```go
type Temp struct{}

func newTempStorage() *Temp {
	return &Temp{}
}

func (s *Temp) Open(f string) (io.ReadWriteCloser, error) {
	return nil, nil
}
```

### 使用

使用工厂方法，用户可以指定想要的存储方式

```go
s, _ := NewStore(data.MemoryStorage)
f, _ := s.Open("file")

n, _ := f.Write([]byte("hello world"))
```