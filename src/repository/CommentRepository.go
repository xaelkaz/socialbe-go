package repository

import (
	"fmt"
	"social-golang/src/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "time"
	"go.mongodb.org/mongo-driver/bson"

	"social-golang/src/models"

)

type CommentRepository struct {
	db *mongo.Collection
	collection string
}

//NewProfileRepositoryMongo
func NewCommentRepository(collection string) *CommentRepository{
	return &CommentRepository{
		db: storage.Comment,
		collection: collection,
	}
}

//Save
func (r *CommentRepository) Save(comment *models.Comment) error{
	res, err := r.db.InsertOne(ctx, comment)
	fmt.Println(res)
	return  err
}

//Update
func (r *CommentRepository) Update(comment *models.Comment) error{
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": comment.ID}, bson.M{"$set": comment})
	fmt.Println(res)
	return err
}

//Delete
func (r *CommentRepository) Delete(id primitive.ObjectID) error{
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *CommentRepository) FindByID(id primitive.ObjectID) (*models.Comment, error){
	var comment models.Comment
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {return nil, err}
	return &comment, nil
}

//FindAll
func (r *CommentRepository) FindAll() (models.Comments, error){
	var comment models.Comments
	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {return nil, err}
	for cursor.Next(ctx) {
		var c models.Comment
		if err := cursor.Decode(&c); err != nil {}
		comment = append(comment, c)
	}
	return comment, nil
}


//FindByName
func (r *CommentRepository) FindByName(name string) (*models.Comment, error){
	var comment models.Comment
	err := r.db.FindOne(ctx, bson.M{"text": name}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

