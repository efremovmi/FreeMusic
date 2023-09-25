package mongodb

import (
	"FreeMusic/internal/app_errors"
	"FreeMusic/internal/models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (m *mongoFileStorage) DownloadFile(ctx context.Context, req models.DownloadFileRequest) (*models.DownloadFileResponse, error) {
	db := m.client.Database(m.databaseName)

	fileInfo, err := m.findIDHexByFileNameAndUserID(ctx, db, req)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile error")
	}

	resp, err := m.getFileStreamByFileIDHex(ctx, db, fileInfo)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile error")
	}

	if err == nil && resp == nil {
		return nil, &app_errors.FileNotFound{
			Message: "file not found",
		}
	}

	return resp, nil
}

func (m *mongoFileStorage) findIDHexByFileNameAndUserID(ctx context.Context, db *mongo.Database, req models.DownloadFileRequest) (*models.FileInfo, error) {
	collection := db.Collection(m.fileCollectionName)
	filter := bson.M{
		"file_name": req.FileName,
		"user_id":   req.UserID,
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.Wrap(err, "findIDHexByFileNameAndUserID: can't get cursor on collections")
	}
	defer cursor.Close(context.Background())

	var fileInfo models.FileInfo
	var isGetDataFromDB bool
	for cursor.Next(context.Background()) {
		isGetDataFromDB = true
		if err := cursor.Decode(&fileInfo); err != nil {
			return nil, errors.Wrap(err, "findIDHexByFileNameAndUserID: can't decode file from db")
		}
		break
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "findIDHexByFileNameAndUserID: cursor error")
	}

	if !isGetDataFromDB {
		return nil, errors.Wrap(err, "findIDHexByFileNameAndUserID: can't get data from db")
	}

	if fileInfo.UserID == 0 {
		return nil, errors.Wrap(err, "findIDHexByFileNameAndUserID: can't bad data from collections 'files'")
	}

	return &fileInfo, nil
}

func (m *mongoFileStorage) getFileStreamByFileIDHex(ctx context.Context, db *mongo.Database, fileInfo *models.FileInfo) (*models.DownloadFileResponse, error) {
	if fileInfo == nil {
		return nil, errors.Wrap(nil, "getFileStreamByFileIDHex: get null fileInfo")
	}

	fs, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, errors.Wrap(err, "getFileStreamByFileIDHex: can't get bucket")
	}

	fileID, err := primitive.ObjectIDFromHex(fileInfo.FileIDHex)
	if err != nil {
		return nil, errors.Wrap(err, "getFileStreamByFileIDHex: can't convert FileIDHex to primitive.ObjectIDFromHex")
	}

	fileStream, err := fs.OpenDownloadStream(fileID)
	if err != nil {
		return nil, errors.Wrap(err, "getFileStreamByFileIDHex: can't get file stream")
	}

	return &models.DownloadFileResponse{
		FileName:   fileInfo.FileName + fileInfo.FileExtension,
		FileStream: fileStream,
	}, nil
}
