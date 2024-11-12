package storage

import "collaboration/types"

type Storage interface {
	Get(int) (*types.User, error)
	Delete(int) error
}
