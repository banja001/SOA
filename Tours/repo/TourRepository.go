package repo

import (
	"context"
	"database-example/model"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TourRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	var tour model.Tour
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := repo.getCollection()

	objID, _ := strconv.Atoi(id)
	err := toursCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&tour)
	if err != nil {
		log.Println("err FindOne: ", err)
		return tour, err
	}

	return tour, nil
}

func (repo *TourRepository) Create(tour *model.Tour) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := repo.getCollection()
	
	id, err := repo.getNextID(ctx)
	if err != nil {
		log.Println("err getNextId: ", err)
		return nil, err
	}
	tour.ID = id

	result, err := toursCollection.InsertOne(ctx, &tour)
	if err != nil {
		log.Println("err InsertOne: ", err)
		return nil, err
	}

	log.Printf("Documents ID: %v\n", result.InsertedID)
	return tour, nil
}


func (repo *TourRepository) Update(tour *model.Tour) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := repo.getCollection()

	filter := bson.M{"_id": tour.ID}
	update := bson.M{
		"$set": bson.M{
			"name":          tour.Name,
			"description":   tour.Description,
			"difficulty":    tour.Difficulty,
			"tags":          tour.Tags,
			"status":        tour.Status,
			"price":         tour.Price,
			"authorId":      tour.AuthorId,
			"equipment":     tour.Equipment,
			"distance":      tour.DistanceInKm,
			"archived":      tour.ArchivedDate,
			"published":     tour.PublishedDate,
			"durations":     tour.Durations,
			"keypoints":     tour.KeyPoints,
			"image":         tour.Image,
		},
	}
	result, err := toursCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	log.Printf("Documents matched: %v\n", result.MatchedCount)
	log.Printf("Documents updated: %v\n", result.ModifiedCount)

	return tour, nil
}

func (repo *TourRepository) FindByAuthorId(authorID string) ([]model.Tour, error) {
	var tours []model.Tour
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := repo.getCollection()

	authorId, _ := strconv.Atoi(authorID)
	filter := bson.M{"authorId": authorId}
	toursCursor, err := toursCollection.Find(ctx, filter)

	if err != nil {
		log.Println("err Find: ", err)
		return nil, err
	}

	err = toursCursor.All(ctx, &tours)
	if err != nil {
		log.Println("err All: ", err)
		return nil, err
	}

	return tours, nil
}

func (repo *TourRepository) getCollection() *mongo.Collection {
	toursDatabase := repo.DatabaseConnection.Database("mongoDemo")
	toursCollection := toursDatabase.Collection("tours")
	return toursCollection
}

func (repo *TourRepository) getNextID(ctx context.Context) (int, error) {
	counterCollection := repo.DatabaseConnection.Database("mongoDemo").Collection("counters")
	filter := bson.M{"_id": "tourID"}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	var result struct {
		Seq int `bson:"seq"`
	}
	err := counterCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.Seq, nil
}