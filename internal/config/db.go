package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenDB(cfg *Config) (*mongo.Client, error) {
	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	opts := options.Client().
		ApplyURI(cfg.Db.DSN).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)).
		SetMaxPoolSize(uint64(cfg.Db.MaxIdleConn)).
		SetMinPoolSize(1).
		SetMaxConnIdleTime(duration)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	ctx, cleanup := context.WithTimeout(context.Background(), 3*time.Second)
	defer cleanup()

	if err := client.Ping(ctx, readpref.Nearest()); err != nil {
		return nil, err
	}

	return client, err
}
