package handler

import (
	Cofig "WEB_SERVER/cofig"
	"WEB_SERVER/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// 	"golang.org/x/mod/modfile"
)

// CREATE

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := Cofig.ConnectTOMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background()) //must write lines till here

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := client.Database("mydatabase").Collection("Users")

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(result)
}

// ALL USER

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	client, err := Cofig.ConnectTOMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background()) //must write lines till here
	collection := client.Database("mydatabase").Collection("Users")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var T_user []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		T_user = append(T_user, user)
	}
	json.NewEncoder(w).Encode(T_user)
}

// USER_BY_ID

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	client, err := Cofig.ConnectTOMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	collection := client.Database("mydatabase").Collection("Users")
	var user models.User
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UPDATE

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	client, err := Cofig.ConnectTOMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer client.Disconnect(context.Background())
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := client.Database("mydatabase").Collection("User")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// DELETE

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	client, err := Cofig.ConnectTOMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("mydatabase").Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}
