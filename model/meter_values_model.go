package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeterValues struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ChargerID     string             `json:"chargerId"`
	ConnectorID   int                `json:"connectorId"`
	TransactionId int                `json:"transactionId"`
	Value         int                `json:"value"`
	Type          string             `json:"type"`
	Unit          string             `json:"unit"`
	CreatedAt     time.Time          `json:"createdAt"`
}
