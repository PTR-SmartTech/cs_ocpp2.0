package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	TransactionId  string             `bson:"transactionId"`
	ChargerID      string             `bson:"chargerId"`
	ConnectorID    int                `bson:"connectorId"`
	StartTimestamp time.Time          `bson:"startTimestamp"`
	StartReason    string             `bson:"startReason"`
	StopTimestamp  time.Time          `bson:"stopTimestamp"`
	StopReason     string             `bson:"stopReason"`
	MeterStart     float64            `bson:"meterStart"`
	MeterStop      float64            `bson:"meterStop"`
	Status         string             `bson:"status"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
