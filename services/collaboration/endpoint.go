package collaboration

import (
	collaboration "collaboration/services/collaboration/service"
	"collaboration/types"
	"context"
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

func MakeConnectEndpoint(svc collaboration.CollaborationService) func(ctx context.Context, client *types.Client) connectResponse {
	return func(ctx context.Context, client *types.Client) connectResponse {
		svc.RegisterClient(client)
		return connectResponse{Success: true}
	}
}

func MakeBroadcastEndpoint(svc collaboration.CollaborationService) func(ctx context.Context, message types.Message) messageResponse {
	return func(ctx context.Context, message types.Message) messageResponse {
		svc.BroadcastMessage(message)
		return messageResponse{Success: true}
	}
}