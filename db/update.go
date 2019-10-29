package db

import (
	"context"
	"time"

	"github.com/letanthang/go_fw/types"
)

func AddStudent(req types.AddStudentReq) (interface{}, error) {
	// counterCol := Client.Database(DBName).Collection("my_sequence")
	// ID := mymongo.GetNextID(counterCol, "student")
	student := types.Student{
		ID:        3,
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
