package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoContext context.Context = context.Background()
	mongoClient  *mongo.Client   = &mongo.Client{}
)

// EstablishDbConnection instantiates a new connection pool and verifies
// that the client can access MongoDB's deployment. If both of these tasks
// conclude error-free, the client is used as an entrypoint to the database.
func EstablishDbConnection(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	newClient, err := mongo.Connect(mongoContext, clientOptions)
	if err != nil {
		return fmt.Errorf("cannot create client: %w", err)
	} else if err = newClient.Ping(mongoContext, nil); err != nil {
		return fmt.Errorf("cannot connect to mongodb: %w", err)
	}

	fmt.Println("mongodb connection successfully established!")
	mongoClient = newClient
	return nil
}

// CloseDbConnection attempts to gracefully disconnect from the database,
// shutting down monitoring goroutines as well as the idle connection pool.
func CloseDbConnection() error {
	if err := mongoClient.Disconnect(mongoContext); err != nil {
		return fmt.Errorf("cannot disconnect from mongodb: %w", err)
	}
	return nil
}
