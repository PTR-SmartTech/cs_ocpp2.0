package model

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connector struct {
	ID          primitive.ObjectID           `bson:"_id,omitempty"`
	ChargerID   string                       `bson:"chargerId"`
	ConnectorID int                          `bson:"connectorId"`
	Status      availability.ConnectorStatus `bson:"status"`
	CreatedAt   time.Time                    `bson:"createdAt"`
	UpdatedAt   time.Time                    `bson:"updatedAt"`
}
