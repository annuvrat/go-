package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoInstance struct {
Client *mongo.Client
   Db  *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

type Employee struct {
ID        string `json:"id" bson:"id"`
FirstName string `json:"first_name" bson:"first_name"`
LastName  string `json:"last_name" bson:"last_name"`
Email     string `json:"email" bson:"email"`
Position  string `json:"position" bson:"position"`
} 


func connect() (*mongo.Client, error) {
    ctx := context.TODO() // or pass a proper context
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
    if err != nil {
        return nil, err
    }
    return client, nil
}

func main() {

	err:= connect()
	if err != nil {
		log.Fatal(err) 
	} 
	 app:=fiber.New()

	 app.Get("/employees",func (c *fiber.Ctx) error{
		return c.SendString("Get all employees")
	 })
	 app.Post("/employees")
	 app.Put("/employees/:id")
	 app.Delete("/employees/:id")

}