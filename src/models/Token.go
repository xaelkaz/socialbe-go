package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID    primitive.ObjectID `json:"userId"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
