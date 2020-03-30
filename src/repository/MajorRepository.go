package repository

import (
	"fmt"
	"social-golang/src/models"
	"social-golang/src/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MajorRepositoryMongo struct {
	db         *mongo.Collection
	collection string
}

//NewProfileRepositoryMongo
func NewMajorRepository(collection string) *MajorRepositoryMongo {
	return &MajorRepositoryMongo{
		db:         storage.Major,
		collection: collection,
	}
}

//Save
func (r *MajorRepositoryMongo) Save(major *models.Major) error {
	res, err := r.db.InsertOne(ctx, major)
	fmt.Println(res)
	return err
}

//Update
func (r *MajorRepositoryMongo) Update(major *models.Major) error {
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": major.ID}, bson.M{"$set": major})
	fmt.Println(res)
	return err
}

//Delete
func (r *MajorRepositoryMongo) Delete(id primitive.ObjectID) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *MajorRepositoryMongo) FindByID(id primitive.ObjectID) (*models.Major, error) {
	var major models.Major
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&major)

	if err != nil {
		return nil, err
	}

	return &major, nil
}

//FindAll
func (r *MajorRepositoryMongo) FindAll() (models.Majors, error) {
	var major models.Majors

	cursor, err := r.db.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var m models.Major
		if err := cursor.Decode(&m); err != nil {
		}
		major = append(major, m)
	}
	return major, nil
}

//FindByName
func (r *MajorRepositoryMongo) FindByName(name string) (*models.Major, error) {
	var major models.Major
	err := r.db.FindOne(ctx, bson.M{"name": name}).Decode(&major)

	if err != nil {
		return nil, err
	}

	return &major, nil
}

func (r *MajorRepositoryMongo) FindByFaculty(facultyName string) (models.Majors, error) {
	var major models.Majors
	cursor, err := r.db.Find(ctx, bson.M{"faculty.name": facultyName})
	if err != nil {

	}
	for cursor.Next(ctx) {
		var m models.Major
		if err := cursor.Decode(&m); err != nil {
		}
		major = append(major, m)
	}
	return major, nil
}

func (r *MajorRepositoryMongo) DeleteByName(name string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"name": name})
	fmt.Println(res)
	return err
}
