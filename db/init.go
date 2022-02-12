package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Client
}

func Init(dataSourceName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(dataSourceName)

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return &MongoDB{db}, nil
}
