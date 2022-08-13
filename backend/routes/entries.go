package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var entryCollection *mongo.Collection = openCollection(Client, "calories")

func AddEntry(c *gin.Context){

}

func GetEntries(c *gin.Context){

}

func GetEntryById(c *gin.Context){

}

func GetEntriesByIngredient(c *gin.Context){

}

func UpdateEntry(c *gin.Context){

}

func UpdateIngredient(c *gin.Context){

}

func DeleteEntry(c *gin.Context){
	entryID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id":docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}