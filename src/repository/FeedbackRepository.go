package repository

import (
	"fmt"
	"social-golang/src/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	// "time"
	"go.mongodb.org/mongo-driver/mongo"

	"social-golang/src/models"
)

type FeedbackRepository struct {
	db         *mongo.Collection
	collection string
}

//NewProfileRepositoryMongo
func NewFeedbackRepository(collection string) *FeedbackRepository {
	return &FeedbackRepository{
		db:         storage.Feedback,
		collection: collection,
	}
}

//Save
func (r *FeedbackRepository) Save(feedback *models.Feedback) error {
	res, err := r.db.InsertOne(ctx, feedback)
	fmt.Println(res)
	return err
}

//Update
func (r *FeedbackRepository) Update(feedback *models.Feedback) error {
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": feedback.ID}, bson.M{"$set": feedback})
	fmt.Println(res)
	return err
}

//Delete
func (r *FeedbackRepository) Delete(feedback *models.Feedback) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": feedback.ID})
	fmt.Println(res)
	return err
}

//FindByID
func (r *FeedbackRepository) FindByID(id primitive.ObjectID) (*models.Feedback, error) {
	var feedback models.Feedback
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&feedback)
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

//FindAll
func (r *FeedbackRepository) FindAll() (models.Feedbacks, error) {
	var feedback models.Feedbacks

	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var f models.Feedback
		if err := cursor.Decode(&f); err != nil {
		}
		feedback = append(feedback, f)
	}

	return feedback, nil
}

//FindByName
func (r *FeedbackRepository) FindByName(name string) (*models.Feedback, error) {
	var feedback models.Feedback
	err := r.db.FindOne(ctx, bson.M{"text": name}).Decode(&feedback)

	if err != nil {
		return nil, err
	}

	return &feedback, nil
}
func (r *FeedbackRepository) FindByCode(code string) (models.Feedbacks, error) {
	var feedback models.Feedbacks
	cursor, err := r.db.Find(ctx, bson.M{"subject.code": code})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var f models.Feedback
		if err := cursor.Decode(&f); err != nil {
		}
		feedback = append(feedback, f)
	}
	return feedback, nil
}
