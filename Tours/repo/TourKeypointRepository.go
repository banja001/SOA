package repo

import (
	"database-example/model"
	"log"
	"strconv"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TourKeypointRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *TourKeypointRepository) FindById(id string) (model.TourKeypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourKeyPointsCollection := repo.getCollection()
	var tourKeypoint model.TourKeypoint
	objID, _ := strconv.Atoi(id)
	err := tourKeyPointsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&tourKeypoint)
	if err != nil {
		log.Println(err)
		return tourKeypoint, err
	}
	return tourKeypoint, nil
}

func (repo *TourKeypointRepository) Create(tourKeypoint *model.TourKeypoint) (model.TourKeypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := repo.getNextID(ctx)
	if err != nil {
		log.Printf("Failed to get next ID: %v", err)
		return model.TourKeypoint{}, err
	}
	tourKeypoint.ID = id

	tourKeyPointsCollection := repo.getCollection()
	_, err = tourKeyPointsCollection.InsertOne(ctx, tourKeypoint)
	if err != nil {
		log.Printf("Failed to insert tourKeypoint: %v", err)
		return model.TourKeypoint{}, err
	}

	log.Println("Tour keypoint created")
	return *tourKeypoint, nil
}

func (repo *TourKeypointRepository) Update(tourKeypoint *model.TourKeypoint) (model.TourKeypoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tourKeyPointsCollection := repo.getCollection()

	filter := bson.M{"_id": tourKeypoint.ID}
	update := bson.M{"$set": bson.M{
		"Name":           tourKeypoint.Name,
		"Description":    tourKeypoint.Description,
		"Image":          tourKeypoint.Image,
		"Latitude":       tourKeypoint.Latitude,
		"Longitude":      tourKeypoint.Longitude,
		"TourID":         tourKeypoint.TourID,
		"Secret":         tourKeypoint.Secret,
		"PositionInTour": tourKeypoint.PositionInTour,
		"PublicPointID":  tourKeypoint.PublicPointID,
	}}

	result, err := tourKeyPointsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to update tourKeypoint: %v", err)
		return model.TourKeypoint{}, err
	}

	log.Printf("Documents matched: %v\n", result.MatchedCount)
	log.Printf("Documents updated: %v\n", result.ModifiedCount)

	return *tourKeypoint, nil
}

func (repo *TourKeypointRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourKeyPointsCollection := repo.getCollection()

	objID, _ := strconv.Atoi(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := tourKeyPointsCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (kpr *TourKeypointRepository) getCollection() *mongo.Collection {
	keyPointDatabase := kpr.DatabaseConnection.Database("mongoDemo")
	keyPointsCollection := keyPointDatabase.Collection("tourKeyPoints")
	return keyPointsCollection
}

func (repo *TourKeypointRepository) getNextID(ctx context.Context) (int, error) {
	counterCollection := repo.DatabaseConnection.Database("mongoDemo").Collection("counters")
	filter := bson.M{"_id": "tourKeypointID"}
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
