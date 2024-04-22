package data

import (
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
