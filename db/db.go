package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/amar-jay/go-graphql/graph/model"
)

type DB struct {
  client *mongo.Client
}

func Connnection() *DB {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil {
    log.Fatal(err)
  }

  return &DB{
    client: client,
  } 
}
func (db *DB) Create(input *model.NewTodo) *model.Todo {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  collection := db.client.Database("gotodos").Collection("todos")
  res, err := collection.InsertOne(ctx, input)
  if err != nil {
    log.Fatal(err)
  }
  return &model.Todo{
    ID: res.InsertedID.(primitive.ObjectID).Hex(),
    Text: input.Text, 
  }
}


func (db *DB) GetByID(ID string) *model.Todo {
  var res = model.Todo{}
  ObjectID, err := primitive.ObjectIDFromHex(ID)
  if err != nil {
    log.Fatal("Invalid ID")
  }
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  collection := db.client.Database("gotodos").Collection("todos")
  //err = collection.FindOne(ctx, bson.M{"id", ObjectID}).Decode(&res)
  err = collection.FindOne(ctx, ObjectID).Decode(&res)
  if err != mongo.ErrNoDocuments { 
    fmt.Println("Record not found.")
  } else if err != nil {
    log.Fatal(err)
  }

  return &res 
}

func (db *DB) GetAll() []*model.Todo {
  var res []*model.Todo
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()
  collection := db.client.Database("gotodos").Collection("todos")
  cur, err := collection.Find(ctx, bson.D{})
  if err != nil {
    log.Fatal(err)
  }
  defer cur.Close(ctx)
  for cur.Next(ctx) {
   var todo *model.Todo
   err := cur.Decode(&todo)
   if err != nil {
     log.Fatal(err)
   }
   res = append(res, todo)
  }

  if err := cur.Err(); err!= nil {
    log.Fatal(err)
  }
  return res
}
