package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Token    string             `json:"token"`
}

var client *mongo.Client
var secret = []byte("my_secret_key")

const USER_ID = "UserId"

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017")
	fmt.Printf("connecting to mongo")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(corsMiddleware)
	r.POST("/login", loginHandler)
	r.GET("/api/card", authMiddleware, getCardsHandler)
	r.POST("/api/card", authMiddleware, postCardsHandler)
	r.DELETE("/api/card", authMiddleware, deleteAllCards)
	r.DELETE("/api/card/:id", authMiddleware, deleteCard)
	r.GET("/api/dog/breed", getBreedsListHandler)

	r.Run(":8080")
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	collection := client.Database("DC-App").Collection("Users")
	result := collection.FindOne(context.Background(), bson.M{"token": tokenString})
	var u User
	err := result.Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			fmt.Println("no auth found")
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking authorization"})
		return
	}

	c.Set(USER_ID, u.Id.Hex())
	c.Next()
}
func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

func loginHandler(c *gin.Context) {
	fmt.Println("Entered Login Handler!")

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		fmt.Println("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	collection := client.Database("DC-App").Collection("Users")
	result := collection.FindOne(context.Background(), bson.M{"username": credentials.Username, "password": credentials.Password})
	var u User
	err := result.Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("found no matchign credentials")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		fmt.Printf("Error Decoding: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking authorization"})
		return
	}
	fmt.Printf("Bound bson %v\n", u)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"id":       u.Id,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Printf("signing token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error authorizing"})
		return
	}

	filter := bson.D{{Key: "_id", Value: u.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: tokenString}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("updating creds: %v\n", err)
		fmt.Printf("updated count %v\n", updateResult.ModifiedCount)
		fmt.Printf("result: %v\n", updateResult)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating authorization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
