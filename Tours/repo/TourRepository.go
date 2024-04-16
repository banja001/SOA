package repo

import (
	"context"
	"database-example/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TourRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	var tour model.Tour
	collection := repo.DatabaseConnection.Database("yourDBName").Collection("tours")
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&tour)
	if err != nil {
		return tour, err
	}
	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (*model.Tour, error) {
	collection := repo.DatabaseConnection.Database("yourDBName").Collection("tours")
	_, err := collection.InsertOne(context.Background(), tour)
	if err != nil {
		return nil, err
	}

	return tour, nil
}

func (repo *TourRepository) FindAll() ([]model.Tour, error) {
	var tours []model.Tour
	collection := repo.DatabaseConnection.Database("yourDBName").Collection("tours")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &tours); err != nil {
		return nil, err
	}
	return tours, nil
}

func (repo *TourRepository) Update(tour *model.Tour) (*model.Tour, error) {
	collection := repo.DatabaseConnection.Database("yourDBName").Collection("tours")
	filter := bson.M{"id": tour.ID}
	update := bson.M{"$set": tour}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return tour, nil
}

func (repo *TourRepository) FindByAuthorId(authorID string) ([]model.Tour, error) {
	var tours []model.Tour
	collection := repo.DatabaseConnection.Database("yourDBName").Collection("tours")
	filter := bson.M{"author_id": authorID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &tours); err != nil {
		return nil, err
	}
	return tours, nil
}
