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

	client, err := connect()
	if err != nil {
		log.Fatal(err) 
	}
	mg.Client = client
	mg.Db = client.Database(dbName)
	 app:=fiber.New()

	 app.Get("/employees",func (c *fiber.Ctx) error{
		query := bson.D{}
		cursor,err:= mg.Db.Collection("employees").Find(c.Context(),query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		var employees []Employee = make([]Employee,0)
		cursor.All(c.Context(),&employees)
		if err := cursor.All(c.Context(),&employees); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(employees)
	 })
	 app.Post("/employees")
	 app.Put("/employees/:id")
	 app.Delete("/employees/:id")

}