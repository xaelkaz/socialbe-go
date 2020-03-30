package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"social-golang/src/models"
	"social-golang/src/storage"
)

type UserRepositoryMongo struct {
	db         *mongo.Collection
	collection string
}

//NewProfileRepositoryMongo
func NewUserRepository(collection string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db:         storage.User,
		collection: collection,
	}
}

//Save
func (r *UserRepositoryMongo) Save(user *models.User) error {
	res, err := r.db.InsertOne(ctx, user)
	fmt.Println(res)
	return err
}

func (r *UserRepositoryMongo) Update(user *models.User) error {
	fmt.Println(user.ID)
	fmt.Println(user.Subject)
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	fmt.Println(res)
	fmt.Println(err)
	return err
}

//Delete
func (r *UserRepositoryMongo) Delete(id primitive.ObjectID) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *UserRepositoryMongo) FindByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

//FindAll
func (r *UserRepositoryMongo) FindAll() (models.Users, error) {
	var user models.Users

	cursor, err := r.db.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
		}
		user = append(user, u)
	}
	return user, nil
}

//FindByName
func (r *UserRepositoryMongo) FindByName(name string) (*models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{"firstname": name}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryMongo) FindByEmail(Email string) (*models.User, error) {
	var user models.User
	err := r.db.FindOne(ctx, bson.M{"email": Email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
