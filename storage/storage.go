package storage

import "collboration/types"

type Storage interface {
	Get(int) *types.User
}
