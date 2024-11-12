package api

import (
	"collaboration/types"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getUserRequest struct {
	ID int `json:"id"`
}

type getUserResponse struct {
	Value *types.User `json:"value"`
	Err   string      `json:"err,omitempty"`
}

type deleteUserRequest struct {
	ID int `json:"id"`
}

type deleteUserResponse struct {
	Err string `json:"err,omitempty"`
}

func makeGetUserByIdEndpoint(svc ApiService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		value, err := svc.HandleGetUserById(req.ID)

		if err != nil {
			return getUserResponse{value, err.Error()}, nil
		}

		return getUserResponse{value, ""}, nil
	}
}

func makeDeleteUserEndpoint(svc ApiService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		err := svc.HandleDeleteUserById(req.ID)

		if err != nil {
			return deleteUserResponse{err.Error()}, err
		}

		return deleteUserResponse{""}, nil
	}
}
