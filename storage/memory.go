package storage

import "collaboration/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "John Memory Doe",
	}, nil
}

func (s *MemoryStorage) Delete(id int) error {
	return nil
}
