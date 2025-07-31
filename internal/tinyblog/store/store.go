package store

import (
	"sync"

	"gorm.io/gorm"
)





type dataStore struct {
	db *gorm.DB
}

var _ IStore = (*dataStore)(nil)

var(
	S *dataStore
	once sync.Once
)

func GetDataStore() *dataStore {
	return S
}

func InitDataStore(db *gorm.DB) {
	once.Do(func(){
		S = &dataStore{db}
	})
}

func (s *dataStore) UserStore() UserStore {
	return newUserStore(s.db)
}

func (s *dataStore) PostStore() PostStore {
	return newPostStore(s.db)
}

func (s *dataStore) DB() *gorm.DB {
	return s.db
}

