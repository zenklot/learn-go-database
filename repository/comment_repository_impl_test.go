package repository

import (
	"context"
	"fmt"
	learn_go_database "learn-go-database"
	"learn-go-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "zez@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 33)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
