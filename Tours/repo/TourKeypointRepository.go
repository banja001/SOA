package repo

import (
	"database-example/model"
	"time"

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TourKeypointRepository struct {
	DatabaseConnection *mongo.Client
}

// func (repo *TourKeypointRepository) FindById(id string) (model.TourKeypoint, error) {
// 	tourKeypoint := model.TourKeypoint{}
// 	dbResult := repo.DatabaseConnection.First(&tourKeypoint, "id = ?", id)
// 	if dbResult != nil {
// 		return tourKeypoint, dbResult.Error
// 	}
// 	return tourKeypoint, nil
// }

func (repo *TourKeypointRepository) FindById(id string) (model.TourKeypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tourKeyPointsCollection := repo.getCollection()

	tourKeypoint := model.TourKeypoint{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := tourKeyPointsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&tourKeypoint)
	if err != nil {
		return tourKeypoint, err
	}
	return tourKeypoint, nil
}

// func (repo *TourKeypointRepository) Create(tourKeypoint *model.TourKeypoint) (model.TourKeypoint, error) {
// 	dbResult := repo.DatabaseConnection.Create(tourKeypoint)
// 	if dbResult.Error != nil {
// 		return *tourKeypoint, dbResult.Error
// 	}
// 	println("Tour keypoints created: ", dbResult.RowsAffected)
// 	return *tourKeypoint, nil
// }

func (repo *TourKeypointRepository) Create(tourKeypoint *model.TourKeypoint) (model.TourKeypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourKeyPointsCollection := repo.getCollection()

	_, err := tourKeyPointsCollection.InsertOne(ctx, &tourKeypoint)
	// if err != nil {
	// 	return err
	// }
	fmt.Printf("Created tourKeypoint:")
	return *tourKeypoint, err
}

func (repo *TourKeypointRepository) Update(tourKeypoint *model.TourKeypoint) (model.TourKeypoint, error) {
	// dbResult := repo.DatabaseConnection.Updates(tourKeypoint)
	// if dbResult.Error != nil {
	// 	return *tourKeypoint, dbResult.Error
	// }
	// println("Tour keypoints updated: ", dbResult.RowsAffected)
	return *tourKeypoint, nil
}

func (repo *TourKeypointRepository) Delete(id string) error {
	// dbResult := repo.DatabaseConnection.Unscoped().Delete(&model.TourKeypoint{}, "id = ?", id)
	// if dbResult.Error != nil {
	// 	return dbResult.Error
	// }
	// println("Tour keypoints deleted: ", dbResult.RowsAffected)
	return nil
}

func (kpr *TourKeypointRepository) getCollection() *mongo.Collection {
	keyPointDatabase := kpr.DatabaseConnection.Database("mongoDemo")
	keyPointsCollection := keyPointDatabase.Collection("tourKeyPoints")
	return keyPointsCollection
}
