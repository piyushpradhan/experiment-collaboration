package api

import (
	"collaboration/storage"
	"collaboration/types"
)

type ApiService interface {
	HandleGetUserById(id int) (*types.User, error) // Returns *types.User and error
	HandleDeleteUserById(id int) error             // Returns only error
}

type apiService struct {
	store storage.Storage
}

func (a *apiService) HandleGetUserById(id int) (*types.User, error) {
	// Retrieve the user by ID
	user, err := a.store.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *apiService) HandleDeleteUserById(id int) error {
	// Delete the user by ID
	err := a.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewApiService() ApiService {
	store := storage.NewMemoryStorage()

	return &apiService{
		store: store,
	}
}
