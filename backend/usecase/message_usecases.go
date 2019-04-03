package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"my-archive/backend/models"

	"github.com/gorilla/websocket"
)

func (uc *usecase) SaveSession(id string, conn *websocket.Conn) error {
	return uc.repo.SaveSession(id, conn)
}

func (uc *usecase) DeleteSession(id string) error {
	return uc.repo.DeleteSession(id)
}

func (uc *usecase) ForwardMessage(b []byte) error {
	message := models.BucketEvent{}
	err := json.Unmarshal(b, &message)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return err
	}
	if len(message.Records) == 0 {
		log.Printf("[ERROR] Could not parse message: %s", string(b))
		return errors.New("")
	}
	c, err := uc.repo.GetSession(message.Records[0].S3.Bucket.OwnerIdentity.PrincipalID)
	if err != nil {
		log.Printf("[ERROR] Could not parse message: %s", err.Error())
		return err
	}
	notification := models.EventNotification{
		Name:        message.Records[0].S3.Object.Key,
		ContentType: message.Records[0].S3.Object.ContentType,
		Size:        message.Records[0].S3.Object.Size,
		Time:        message.Records[0].EventTime,
	}
	return c.WriteJSON(&notification)
}
