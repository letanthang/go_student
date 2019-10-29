package types

import (
	"time"
)

type Student struct {
	ID        int       `bson:"id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Age       int       `bson:"age"`
	ClassName string    `bson:"class"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type AddStudentReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassName string `json:"class_name"`
	Email     string `json:"email" validate:"email,min=4"`
}
