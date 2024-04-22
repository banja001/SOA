package repo

import (
	"Followers/model"
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepo struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func New(logger *log.Logger) (*FollowerRepo, error) {
	uri := "bolt://localhost/7687/neo4j"
	user := "neo4j"
	pass := "12345678"
	// uri := os.Getenv("NEO4J_DB")
	// user := os.Getenv("NEO4J_USERNAME")
	// pass := os.Getenv("NEO4J_PASS")
    auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	return &FollowerRepo{
		driver: driver,
		logger: logger,
	}, nil
}

func (fr *FollowerRepo) CheckConnection() {
	ctx := context.Background()
	err := fr.driver.VerifyConnectivity(ctx)
	if err != nil {
		fr.logger.Panic(err)
		return
	}

	fr.logger.Printf(`Neo4J server adress: %s`, fr.driver.Target().Host)
}

func (fr *FollowerRepo) CloseDriverConnection(ctx context.Context) {
	fr.driver.Close(ctx)
}

// GetAllNodesWithFollowerLabel retrieves all nodes labeled as Follower from the Neo4j database.
func (fr *FollowerRepo) GetAllFollowerNodes() (model.Followers, error) {
    ctx := context.Background()
    session := fr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
    defer session.Close(ctx)

    // ExecuteRead for read transactions (Read and queries)
    followerResults, err := session.ExecuteRead(ctx,
        func(transaction neo4j.ManagedTransaction) (interface{}, error) {
            result, err := transaction.Run(ctx,
                `MATCH (follower:Follower)
                RETURN follower.FollowerId as FollowerId, follower.FollowingId as FollowingId`,
                map[string]interface{}{})
            if err != nil {
                return nil, err
            }

            var followers model.Followers
            for result.Next(ctx) {
                record := result.Record()
                followerId, ok := record.Get("FollowerId")
                if !ok || followerId == nil {
                    followerId = 0
                }
                followingId, ok := record.Get("FollowingId")
                if !ok || followingId == nil {
                    followingId = 0
                }

                follower := &model.Follower{
                    FollowerId:   followerId.(int64),
                    FollowingId:  followingId.(int64),
                    Notification: model.FollowerNotification{}, // Assuming there's no notification data in this query
                }
                followers = append(followers, follower)
            }
            return followers, nil
        })
    if err != nil {
        fr.logger.Println("Error querying followers:", err)
        return nil, err
    }
    return followerResults.(model.Followers), nil
}



