package repository

import (
	"context"
	"log"
	"time"

	"github.com/dogg5432/cs_ocpp2.0/database"
	"github.com/dogg5432/cs_ocpp2.0/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionsRepository() TransactionRepository {
	return &transactionRepository{
		collection: database.Client.Collection("transactions"),
	}
}

func (c transactionRepository) Create(transaction *model.Transaction) error {
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	var _, err = c.collection.InsertOne(context.TODO(), transaction)
	if err != nil {
		return err
	}
	return nil
}

func (c transactionRepository) FindOne(transactionID string) (model.Transaction, error) {
	var transaction model.Transaction
	filter := bson.M{"transactionId": transactionID}
	result := c.collection.FindOne(context.TODO(), filter)
	err := result.Decode(&transaction)
	log.Print(filter)
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (c transactionRepository) Update(transaction *model.Transaction) error {
	var filter = bson.M{"transactionId": transaction.TransactionId}
	transaction.UpdatedAt = time.Now()
	var transactionUpdated = bson.M{"$set": transaction}
	var _, err = c.collection.UpdateOne(context.TODO(), filter, transactionUpdated)
	if err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		panic(err)
	}
	return nil
}