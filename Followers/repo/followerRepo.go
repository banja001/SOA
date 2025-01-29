package repo

import (
	"Followers/model"
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type FollowerRepo struct {
	driver neo4j.DriverWithContext
	logger *log.Logger
}

func New(logger *log.Logger) (*FollowerRepo, error) {
	//uri := "bolt://localhost/7666/neo4j"
	//user := "neo4j"
	//pass := "12345678"
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")
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

func (fr *FollowerRepo) GetAllPersonsNodes() (*model.Persons, error) {
	ctx := context.Background()
	session := fr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	personResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				`MATCH (person:Person)
                 RETURN person`,
				map[string]interface{}{})
			if err != nil {
				return nil, err
			}

			var persons model.Persons
			for result.Next(ctx) {
				record := result.Record()
				personNode, ok := record.Get("person")
				if !ok {
					continue
				}
				person, ok := personNode.(neo4j.Node) // Check and convert to neo4j.Node
				if !ok {
					continue
				}

				modelPerson, err := convertNodeToModelPerson(person)
				if err != nil {
					return nil, err // Handle error if conversion fails
				}
				persons = append(persons, modelPerson)
			}
			if err := result.Err(); err != nil {
				return nil, err
			}

			return &persons, nil
		})
	if err != nil {
		fr.logger.Println("Error querying persons:", err)
		return nil, err
	}

	return personResults.(*model.Persons), nil
}

func convertNodeToModelPerson(node neo4j.Node) (*model.Person, error) {
	var p model.Person
	var ok bool

	// Extract fields from the node and check for existence and correct type
	if p.ID, ok = node.Props["id"].(int64); !ok {
		p.ID = 0
	}
	if p.Name, ok = node.Props["name"].(string); !ok {
		p.Name = "" // Set default if not present or wrong type
	}
	if p.Surname, ok = node.Props["surname"].(string); !ok {
		p.Surname = "" // Set default if not present or wrong type
	}
	if p.Email, ok = node.Props["email"].(string); !ok {
		p.Email = "" // Set default if not present or wrong type
	}

	return &p, nil
}

func (fr *FollowerRepo) RewriteFollower(updatedFollower *model.Follower) error {
	ctx := context.Background()
	session := fr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	fol, _ := fr.IsFollowed(int(updatedFollower.FollowerId), int(updatedFollower.FollowedId))
	fr.logger.Println(fol)
	if !fol {
		_, err := session.ExecuteWrite(ctx,
			func(transaction neo4j.ManagedTransaction) (any, error) {
				result, err := transaction.Run(ctx,
					"MATCH (a:Person), (b:Person) "+
						"WHERE a.id = $id1 AND b.id = $id2 "+
						"CREATE (a)-[r:FOLLOWS]->(b) "+
						"RETURN a,b",
					map[string]any{
						"id1": updatedFollower.FollowerId,
						"id2": updatedFollower.FollowedId,
					})
				if err != nil {
					println("Problem\n")
					return nil, err
				}

				if result.Next(ctx) {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err != nil {
			fr.logger.Println("Error updating follower:", err)
			return err
		}
	}
	fr.logger.Println("Relationship already exists!")
	return nil
}

func (fr *FollowerRepo) GetAllRecomended(id int, uid int) (*model.Persons, error) {
	ctx := context.Background()
	session := fr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	fr.logger.Println(id)
	personResults, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (interface{}, error) {
			result, err := transaction.Run(ctx,
				`MATCH (p:Person {id: $id})-[:FOLLOWS]->(following)
				RETURN following`,
				map[string]interface{}{
					"id": id,
				})
			if err != nil {
				return nil, err
			}

			var persons model.Persons
			for result.Next(ctx) {
				record := result.Record()
				personNode, ok := record.Get("following")
				if !ok {
					continue
				}
				person, ok := personNode.(neo4j.Node) // Check and convert to neo4j.Node
				if !ok {
					continue
				}
				modelPerson, err := convertNodeToModelPerson(person)
				if err != nil {
					return nil, err // Handle error if conversion fails
				}
				fol, _ := fr.IsFollowed(uid, int(modelPerson.ID))
				fr.logger.Println(fol)
				fr.logger.Println(modelPerson.ID)
				if !fol {
					persons = append(persons, modelPerson)
				}

			}
			if err := result.Err(); err != nil {
				return nil, err
			}

			return &persons, nil
		})
	if err != nil {
		fr.logger.Println("Error querying persons:", err)
		return nil, err
	}

	return personResults.(*model.Persons), nil
}

func (fr *FollowerRepo) IsFollowed(id1 int, id2 int) (bool, error) {
	ctx := context.Background()
	session := fr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	ok, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (p1:Person {id: $id2})-[:FOLLOWS]->(p2:Person {id: $id1}) "+
					"RETURN p1, p2",
				map[string]any{
					"id1": id1,
					"id2": id2,
				})
			if err != nil {
				println("Problem\n")
				return nil, err
			}
			fr.logger.Println("ID1=$id1\nID2=$id2", id1, id2)
			if result.Next(ctx) {
				return true, nil
			}

			return false, nil

		})
	if err != nil {
		fr.logger.Println("Error updating follower:", err)
		return false, err
	}
	return ok.(bool), nil
}
