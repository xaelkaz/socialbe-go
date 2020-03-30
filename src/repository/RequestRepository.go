package repository

import (
	"fmt"
	"social-golang/src/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"social-golang/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

type RequestRepository struct {
	db         *mongo.Collection
	collection string
}

//NewProfileRepositoryMongo
func NewRequestRepository(collection string) *RequestRepository {
	return &RequestRepository{
		db:         storage.Request,
		collection: collection,
	}
}

//Save
func (r *RequestRepository) Save(request *models.Request) error {
	res, err := r.db.InsertOne(ctx, request)
	fmt.Println(res)
	return err
}

//Update
func (r *RequestRepository) Update(request *models.Request) error {
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": request.ID}, bson.M{"$set": request})
	fmt.Println(res)
	return err
}

//Delete
func (r *RequestRepository) Delete(request *models.Request) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": request.ID})
	fmt.Println(res)
	return err
}

//FindByID
func (r *RequestRepository) FindByID(id primitive.ObjectID) (*models.Request, error) {
	var request models.Request
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}

//FindAll
func (r *RequestRepository) FindAll() (models.Requests, error) {
	var request models.Requests

	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var r models.Request
		if err := cursor.Decode(&r); err != nil {
		}
		request = append(request, r)
	}
	return request, nil
}

//FindByName
func (r *RequestRepository) FindByName(name string) (*models.Request, error) {
	var request models.Request
	err := r.db.FindOne(ctx, bson.M{"text": name}).Decode(&request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}
func (r *RequestRepository) FindByCode(code string) (models.Requests, error) {
	var request models.Requests
	cursor, err := r.db.Find(ctx, bson.M{"subject.code": code})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var r models.Request
		if err := cursor.Decode(&request); err != nil {
		}
		request = append(request, r)
	}
	return request, nil
}
