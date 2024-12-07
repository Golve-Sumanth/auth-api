package utils

import (
	"context"
	"fmt"
	"log"

	"auth-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection
var revokedTokensCollection *mongo.Collection

func InitializeMongoDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userCollection = client.Database("authdb").Collection("users")
	revokedTokensCollection = client.Database("authdb").Collection("revoked_tokens")
}

func AddUser(email string, password string) (*models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
	}
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func RevokeToken(token string) error {
	revokedToken := models.RevokedToken{
		Token: token,
	}
	// log.Println("Revoking token:", token)

	_, err := revokedTokensCollection.InsertOne(context.Background(), revokedToken)
	if err != nil {
		// log.Println("Failed to revoke token:", err)
		return fmt.Errorf("failed to insert revoked token into database: %v", err)
	}
	// log.Println("Token revoked successfully:", token)
	return nil
}

func CheckIfTokenRevoked(token string) bool {
	var revokedToken models.RevokedToken
	err := revokedTokensCollection.FindOne(context.Background(), bson.M{"token": token}).Decode(&revokedToken)
	return err == nil
}
