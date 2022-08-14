package routes

import (
	"context"
	"fmt"
	"github.com/AshiishKarhade/calorie-tracker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

var validate = validator.New()

func AddEntry(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	var entry models.Entry

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	//validationErr := validate.Struct(entry)
	//if validationErr != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
	//	fmt.Println(validationErr)
	//	return
	//}
	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order item was not created"})
		fmt.Println("Order item was not created")
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

func GetEntries(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{}) //bson.M{} - empty gets all documents

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntryById(c *gin.Context){
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry bson.M
	err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)
	defer cancel()
}

//func GetEntriesByIngredient(c *gin.Context){
//
//}

func UpdateEntry(c *gin.Context){
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(entry)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, updateErr := entryCollection.ReplaceOne(ctx,
							bson.M{"_id": docID},
							bson.M{
								"dish": entry.Dish,
								"calories": entry.Calories,
								"protein": entry.Protein,
								"fat" : entry.Fat,
								"ingredients": entry.Ingredients,
							})
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		fmt.Println("Order item was not updated")
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

//func UpdateIngredient(c *gin.Context){
//
//}

func DeleteEntry(c *gin.Context){
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id":docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}