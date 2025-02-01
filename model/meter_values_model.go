package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeterValues struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ChargerID     string             `bson:"chargerId"`
	ConnectorID   int                `bson:"connectorId"`
	TransactionId int                `bson:"transactionId"`
	Value         int                `bson:"value"`
	Type          string             `bson:"type"`
	Unit          string             `bson:"unit"`
	CreatedAt     time.Time          `bson:"createdAt"`
}
