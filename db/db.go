package db

import (
_    "context"
_    "time"

    "go.mongodb.org/mongo-driver/mongo"
_    "go.mongodb.org/mongo-driver/mongo/options"
_    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
  client *mongo.Client
}
