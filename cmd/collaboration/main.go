package collaboration

import (
	service "collaboration/services/collaboration/service"
	transport "collaboration/services/collaboration/transport"
)

func InitializeCollaboration() {
	collaborationService := service.NewCollaborationService()

	go collaborationService.Run()

	transport.StartWebSocketServer(collaborationService, ":7071")
	
	// Block unil an error occurs
	select {}
}