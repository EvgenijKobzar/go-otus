package memory

import (
	"otus/internal/model/catalog"
	"sync"
)

type Repository[T catalog.HasId] struct {
	imx    sync.RWMutex
	items  map[int]T
	nextId int
}
