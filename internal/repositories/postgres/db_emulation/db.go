package db_emulation

type DbEmulation[T any] map[int]T

type SetIdFuncType[T any] func(obj *T, id int)

func NewDbEmulation[T any]() DbEmulation[T] {
	db := make(DbEmulation[T])
	return db
}

func (db *DbEmulation[T]) Create(obj *T, SetIdFunc SetIdFuncType[T]) *T {
	newId := db.getNewId()

	(*db)[newId] = *obj
	SetIdFunc(obj, newId)

	return obj
}

func (db *DbEmulation[T]) GetAll() []*T {
	out := make([]*T, 0, len(*db))

	for _, obj := range *db {
		out = append(out, &obj)
	}

	return out
}

func (db *DbEmulation[T]) GetById(id int) *T {
	obj, exists := (*db)[id]
	if !exists {
		return nil
	}

	return &obj
}

func (db *DbEmulation[T]) getNewId() int {
	return db.getMaxId() + 1
}

func (db *DbEmulation[T]) getMaxId() int {
	maxId := 0

	for id := range *db {
		maxId = max(maxId, id)
	}

	return maxId
}

func (db *DbEmulation[T]) addDummyData(objs []T, SetIdFunc SetIdFuncType[T]) *DbEmulation[T] {
	for _, obj := range objs {
		db.Create(&obj, SetIdFunc)
	}
	return db
}
