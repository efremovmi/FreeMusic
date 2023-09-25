package mongodb

import (
	"FreeMusic/internal/config"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoFileStorage struct {
	client             *mongo.Client
	databaseName       string
	fileCollectionName string
}

func NewMongoFileStorage(config config.Config) (*mongoFileStorage, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
		config.DbFilesUsername, config.DbFilesPassword,
		config.DbFilesHost, config.DbFilesPort))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "NewMongoFileStorage error")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewMongoFileStorage error")
	}

	return &mongoFileStorage{
		client:             client,
		databaseName:       config.DbFilesName,
		fileCollectionName: config.DBFileCollectionName,
	}, nil
}

func (m *mongoFileStorage) Disconnect(ctx context.Context) {
	m.client.Disconnect(ctx)
}
