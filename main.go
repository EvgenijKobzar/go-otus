package main

import "fmt"

type Serial struct {
	Id               int     `json:"id,omitempty"`
	Title            string  `json:"title,omitempty"`
	FileId           int     `json:"file_id,omitempty"`
	Description      string  `json:"description,omitempty"`
	Rating           float64 `json:"rating,omitempty"`
	Duration         float64 `json:"duration,omitempty"`
	Sort             int     `json:"sort,omitempty"`
	ProductionPeriod string  `json:"production_period,omitempty"`
	Quality          string  `json:"quality,omitempty"`
}

func (s Serial) GetId() int {
	return s.Id
}

func (s Serial) SetId(id int) {
	s.Id = int(id)
}

func NewSerial() Serial {
	return Serial{}
}

type HasId interface {
	GetId() int
	SetId(id int)
}

type Repository[T HasId] struct {
	items  map[int]T
	nextId int
}

func NewRepository[T HasId]() *Repository[T] {
	return &Repository[T]{
		nextId: 1,
		items:  make(map[int]T),
	}
}

type Usecase[T HasId] struct {
	repo *Repository[T]
}

func NewUsecase[T HasId](repo *Repository[T]) *Usecase[T] {
	return &Usecase[T]{
		repo: repo,
	}
}

func (uc *Usecase[T]) Create() (T, error) {
	var entity T
	switch any(entity).(type) {
	case *Serial:
		serial := NewSerial()
		fmt.Println(serial)
		entity = any(serial).(T)
		fmt.Println("123")
	}

	// Сохраняем через репозиторий
	err := uc.repo.Save(entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r *Repository[T]) Save(entity T) error {
	if entity.GetId() == 0 {
		entity.SetId(r.nextId)
		r.nextId++
	}
	r.items[entity.GetId()] = entity
	return nil
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var items []T
	for _, entity := range r.items {
		items = append(items, entity)
	}
	return items, nil
}

func main() {
	repo := NewRepository[Serial]()
	NewUsecase(repo).Create()
	fmt.Println(repo.GetAll())
}
