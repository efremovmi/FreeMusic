package mongodb

import (
	"FreeMusic/internal/models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (m *mongoFileStorage) DropFile(ctx context.Context, req models.DropFileRequest) error {
	db := m.client.Database(m.databaseName)

	fileIDHex, err := m.dropFileInfo(db, req)
	if err != nil {
		return errors.Wrap(err, "DropFile error")
	}

	err = m.dropFileInFileStorage(ctx, db, fileIDHex)
	if err != nil {
		return errors.Wrap(err, "DropFile error")
	}

	return nil
}

func (m *mongoFileStorage) dropFileInfo(db *mongo.Database, req models.DropFileRequest) (string, error) {
	filter := bson.M{
		"file_name": req.FileName,
		"user_id":   req.UserID,
	}
	var deletedDocument models.FileInfo

	collection := db.Collection(m.fileCollectionName)

	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedDocument)
	if err != nil {
		return "", errors.Wrap(err, "dropFileInfo: can't drop file info")
	}

	return deletedDocument.FileIDHex, nil
}

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
