package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entry is a row of Specific Food
type Entry struct {
	ID			primitive.ObjectID `bson:"id"`
	Dish		*string				`json:"dish"`
	Ingredients *string				`json:"ingredients"`
	Calories	*float64			`json:"calories"`
	Protein		*float64			`json:"protein"`
	Fat			*float64			`json:"fat"`
}