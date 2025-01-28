package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	TransactionId  int                `json:"transactionId"`
	ChargerID      string             `json:"chargerId"`
	ConnectorID    int                `json:"connectorId"`
	StartTimestamp time.Time          `json:"startTimestamp"`
	StopTimestamp  time.Time          `json:"stopTimestamp"`
	MeterStart     int                `json:"meterStart"`
	MeterStop      int                `json:"meterStop"`
	Status         string             `json:"status"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}
