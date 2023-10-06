package mongodb

import (
	appError "FreeMusic/internal/app_errors"
	"FreeMusic/internal/models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoFileStorage) GetAllMusicFilesInfo(ctx context.Context, userID uint64) (*models.GetAllMusicFilesInfoResponse, error) {
	db := m.client.Database(m.databaseName)

	collection := db.Collection(m.fileCollectionName)
	filter := bson.M{
		"user_id": userID,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllMusicFilesInfoResponse: can't get cursor on collections")
	}
	defer cursor.Close(context.Background())

	var fileInfo models.FileInfo
	var resp models.GetAllMusicFilesInfoResponse
	var isGetDataFromDB bool

	resp.InfoAboutMusicFile = make([]models.InfoAboutMusicFile, 0)
	for cursor.Next(context.Background()) {
		isGetDataFromDB = true
		if err := cursor.Decode(&fileInfo); err != nil {
			return nil, errors.Wrap(err, "GetAllMusicFilesInfo: can't decode file from db")
		}

		resp.InfoAboutMusicFile = append(resp.InfoAboutMusicFile, models.InfoAboutMusicFile{
			Artist:   fileInfo.Artist,
			FileName: fileInfo.FileName,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "GetAllMusicFilesInfo: cursor error")
	}

	if !isGetDataFromDB {
		return nil, errors.Wrap(&appError.NoData{Message: "no data"}, "GetAllMusicFilesInfo: can't get data from db")
	}

	return &resp, nil
}
