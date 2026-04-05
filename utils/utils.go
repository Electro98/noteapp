package utils

type ServerMessage struct {
	Message string `json:"Message"`
}

func JsonMessage(message string) ServerMessage {
	return ServerMessage{Message: message}
}
