package db

import (
	"gallery/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserByEmail(email string) (models.User, error) {
	col := dbInstance.Collection("users")
	var u models.User
	err := col.FindOne(_ctx(), bson.M{"email": email}).Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func AllUsers() ([]models.User, error) {
	col := dbInstance.Collection("users")
	var users []models.User
	v, err := col.Find(_ctx(), bson.M{})
	if err != nil {
		return users, err
	}
	for v.Next(_ctx()) {
		var u models.User
		err = v.Decode(&u)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, err

}

func UserByID(id string) (models.User, error) {
	col := dbInstance.Collection("users")
	var u models.User
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return u, err
	}
	err = col.FindOne(_ctx(), bson.M{"_id": _id}).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
