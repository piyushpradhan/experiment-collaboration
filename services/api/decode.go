package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserRequest

	// Check if the request body is empty
	if r.Body == nil || r.ContentLength == 0 {
		return request, nil // Return the empty request
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeDeleleUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
