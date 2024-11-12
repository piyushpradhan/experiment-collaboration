package collaboration

import (
	collaboration "collaboration/services/collaboration/service"
	"collaboration/types"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type connectRequest struct {
	Client *types.Client
}

type connectResponse struct {
	Success bool `json:"success"`
}

type messageRequest struct {
	Message types.Message
}

type messageResponse struct {
	Success bool `json:"success"`
}

func MakeConnectEndpoint(svc collaboration.CollaborationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(connectRequest)
		svc.RegisterClient(req.Client)
		return connectResponse{Success: true}, nil
	}
}

func MakeBroadcastEndpoint(svc collaboration.CollaborationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(messageRequest)
		svc.BroadcastMessage(req.Message)
		return messageResponse{Success: true}, nil
	}
}
