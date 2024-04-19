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

type SessionRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *SessionRepository) FindById(id string) (model.Session, error) {
	session := model.Session{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sessionsCollection := repo.getCollection()

	objID, _ := strconv.Atoi(id)

	err := sessionsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&session)
	if err != nil {
		log.Println("err FindOne: ", err)
		return session, err
	}

	return session, nil
}

func (repo *SessionRepository) Create(session *model.Session) (*model.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sessionsCollection := repo.getCollection()

	id, err := repo.getNextID(ctx)
	if err != nil {
		log.Println("err getNextId: ", err)
		return nil, err
	}
	session.ID = id

	result, err := sessionsCollection.InsertOne(ctx, &session)
	if err != nil {
		log.Println("err InsertOne: ", err)
		return nil, err
	}

	log.Printf("Documents ID: %v\n", result.InsertedID)

	return session, nil
}

func (repo *SessionRepository) Update(session *model.Session) (*model.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sessionsCollection := repo.getCollection()

	filter := bson.M{"_id": session.ID}
	update := bson.M{
		"$set": bson.M{
			"tourId":                 session.TourId,
			"touristId":              session.TouristId,
			"locationId":             session.LocationId,
			"status":                 session.SessionStatus,
			"transportation":         session.Transportation,
			"distanceCrossed":        session.DistanceCrossedPercent,
			"lastActivity":           session.LastActivity,
			"completedKeyPoints":     session.CompletedKeyPoints,
		},
	}

	result, err := sessionsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	log.Printf("Documents matched: %v\n", result.MatchedCount)
	log.Printf("Documents updated: %v\n", result.ModifiedCount)
	
	return session, nil
}

func (repo *SessionRepository) getCollection() *mongo.Collection {
	toursDatabase := repo.DatabaseConnection.Database("mongoDemo")
	toursCollection := toursDatabase.Collection("sessions")
	return toursCollection
}

func (repo *SessionRepository) getNextID(ctx context.Context) (int, error) {
	counterCollection := repo.DatabaseConnection.Database("mongoDemo").Collection("counters")
	filter := bson.M{"_id": "sessionID"}
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