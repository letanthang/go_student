package db

import (
	"context"
	"time"

	"github.com/letanthang/go_student/mymongo"
	"github.com/letanthang/go_student/types"
)

func AddStudent(req types.AddStudentReq) (interface{}, error) {
	counterCol := Client.Database(DBName).Collection("my_sequence")
	id, err := mymongo.GetNextID(counterCol, "student_seq")
	if err != nil {
		return nil, err
	}
	student := types.Student{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		ClassName: req.ClassName,
		Email:     req.Email,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	res, err := Client.Database(DBName).Collection(Col).InsertOne(context.TODO(), student)

	if err != nil {
		return nil, err
	}

	return res, nil
}
