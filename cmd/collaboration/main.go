package collaboration

import (
	"collaboration/services/collaboration"
	service "collaboration/services/collaboration/service"
)

func InitializeCollaboration() {
	collaborationService := service.NewCollaborationService()

	go collaborationService.Run()

	collaboration.StartWebSocketServer(collaborationService, ":7071")

	// Block unil an error occurs
	select {}
}

