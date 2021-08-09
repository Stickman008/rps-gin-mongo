package signup

import (
	"log"
	"rock-paper-scissor/app/users"
	"rock-paper-scissor/config"
	"rock-paper-scissor/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToDatabase(user *users.User) (primitive.ObjectID, error) {
	client, dbctx, cancel := database.GetConnection()
	defer cancel()
	defer client.Disconnect(dbctx)
	// task.ID = primitive.NewObjectID()

	result, err := client.Database(config.DB.Name).Collection("tasks").InsertOne(dbctx, user)
	if err != nil {
		log.Printf("Could not create Task: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}
