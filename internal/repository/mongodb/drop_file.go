package mongodb

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (m *mongoFileStorage) dropFileInFileStorage(ctx context.Context, db *mongo.Database, fileIDHex string) error {
	fs, err := gridfs.NewBucket(db)
	if err != nil {
		return errors.Wrap(err, "dropFileInFileStorage: can't get bucket")
	}

	fileID, err := primitive.ObjectIDFromHex(fileIDHex)
	if err != nil {
		return errors.Wrap(err, "dropFileInFileStorage: can't convert FileIDHex to primitive.ObjectIDFromHex")
	}

	err = fs.Delete(fileID)
	if err != nil {
		return errors.Wrap(err, "dropFileInFileStorage: can't delete file by fileIDHex")
	}

	return nil
}
