package wpp

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type Notification struct {
	Number  string
	Message string
	Server  string
}

func EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("event:", v)
	}
}

func CheckNumber(number string) bool {
	if len(number) != 11 {
		return false
	}
	return true
}

func MakeNotification(number string, message string, isGroup bool) Notification {
	if isGroup {
		return Notification{
			Number:  GetGroup(number),
			Message: message,
			Server:  types.GroupServer,
		}
	}
	if CheckNumber(number) {
		return Notification{
			Number:  fmt.Sprintf("55%s", number),
			Message: message,
			Server:  types.DefaultUserServer,
		}
	}
	return Notification{}
}

func MakeJid(destination string, server string) types.JID {
	return types.NewJID(destination, server)
}

func MakeTextMessage(message string) *waE2E.Message {
	return &waE2E.Message{Conversation: proto.String(message)}
}

func SendMessage(notification Notification) error {
	fmt.Println(notification)
	_, err := ConnClient.SendMessage(
		context.Background(),
		MakeJid(notification.Number, notification.Server),
		MakeTextMessage(notification.Message),
	)
	if err != nil {
		return err
	}
	return nil
}
