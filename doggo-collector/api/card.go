package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Breed string             `json:"breed"`
	Photo string             `json:"photo"`
}

// type CardService interface {
// 	// GetCards(userId string) ([]Card, error)
// 	// CreateCard(c Card) error
// 	// DeleteCard(cardId string) error
// 	// DeleteAllCards(userId string) error
// }
// type cardService struct {
// 	database Database
// }

//	func newCardService(db Database) CardService {
//		return cardService{
//			database: db,
//		}
//	}
func getCardsHandler(c *gin.Context) {
	userID := c.GetString(USER_ID)
	collection := client.Database("Cards").Collection(userID)
	var cards []Card
	response := gin.H{"cards": cards}

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Printf("Error collection find: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching cards"})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var card Card
		err := cur.Decode(&card)
		if err != nil {
			fmt.Printf("Error parse: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding cards"})
			return
		}
		cards = append(cards, card)
	}

	if len(cards) > 0 {
		response["cards"] = cards
	}

	c.JSON(http.StatusOK, response)
}
func postCardsHandler(c *gin.Context) {
	userID := c.GetString(USER_ID)
	var request struct {
		Label string `json:"breedLabel"`
		Path  string `json:"breedPath"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	photo, err := getPhotoForBreed(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error making breed card"})
		return
	}
	var card = Card{
		Breed: request.Label,
		Photo: photo,
	}

	collection := client.Database("Cards").Collection(userID)
	_, err = collection.InsertOne(context.Background(), card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting card to DB"})
	}

	c.JSON(http.StatusOK, gin.H{"card": card})
}
func deleteAllCards(c *gin.Context) {
	userID := c.GetString(USER_ID)
	collection := client.Database("Cards").Collection(userID)

	result, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting card to DB"})
	}
	fmt.Printf("Deleted %d", result.DeletedCount)
	c.JSON(http.StatusOK, gin.H{"deleted": result.DeletedCount})
}

func deleteCard(c *gin.Context) {
	userID := c.GetString(USER_ID)
	cardId := c.Param("id")
	fmt.Printf("Card ID: %v\n", cardId)
	collection := client.Database("Cards").Collection(userID)
	objId, _ := primitive.ObjectIDFromHex(cardId)
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "problem deleting card",
		})
	}
	c.JSON(http.StatusOK, gin.H{"deleted": result.DeletedCount})
}
