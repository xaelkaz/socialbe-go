package repository

import (
	"fmt"
	"social-golang/src/models"
	"social-golang/src/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	//"go.mongodb.org/storage-driver/storage/options"
	"context"
	//"log"
)

//profileRepositoryMongo
type FacultyRepositoryMongo struct {
	db         *mongo.Collection
	collection string
}

var ctx = context.Background()

//NewProfileRepositoryMongo
func NewFacultyRepositoryMongo(collection string) *FacultyRepositoryMongo {
	return &FacultyRepositoryMongo{
		db:         storage.Faculty,
		collection: collection,
	}
}

//Save
func (r *FacultyRepositoryMongo) Save(faculty *models.Faculty) error {
	res, err := r.db.InsertOne(ctx, faculty)
	fmt.Println(res)
	return err
}

//Update
func (r *FacultyRepositoryMongo) Update(faculty *models.Faculty) error {
	res, err := r.db.UpdateOne(ctx, bson.M{"_id": faculty.ID}, bson.M{"$set": faculty})
	fmt.Println(res)
	return err
}

//Delete
func (r *FacultyRepositoryMongo) Delete(id primitive.ObjectID) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *FacultyRepositoryMongo) FindByID(id primitive.ObjectID) (*models.Faculty, error) {
	var faculty models.Faculty
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&faculty)
	if err != nil {
		return nil, err
	}
	return &faculty, nil
}

//FindAll
func (r *FacultyRepositoryMongo) FindAll() (models.Faculties, error) {
	var faculty models.Faculties
	cursor, err := r.db.Find(ctx, bson.M{})

	if err != nil {
	}
	for cursor.Next(ctx) {
		var f models.Faculty
		err := cursor.Decode(&f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(f)
		faculty = append(faculty, f)
	}
	fmt.Println(faculty)
	return faculty, nil
}

func (r *FacultyRepositoryMongo) FindByName(name string) (*models.Faculty, error) {
	var faculty models.Faculty
	err := r.db.FindOne(ctx, bson.M{"name": name}).Decode(&faculty)

	if err != nil {
		return nil, err
	}

	return &faculty, nil
}
func (r *FacultyRepositoryMongo) DeleteByName(name string) error {
	res, err := r.db.DeleteOne(ctx, bson.M{"name": name})
	fmt.Println(res)
	return err
}
