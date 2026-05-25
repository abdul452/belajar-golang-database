package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/abdul452/belajar-golang-database/connection"
	"github.com/abdul452/belajar-golang-database/entity"
	"github.com/abdul452/belajar-golang-database/repository"
)

func TestCommentInsert(t *testing.T) {
	db := connection.GetConnection()
	defer db.Close()

	commentRepository := repository.NewCommentRepository(db)
	ctx := context.Background()

	comment := entity.Comments{
		Email:   "test@example.com",
		Comment: "This is a test comment",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	db := connection.GetConnection()
	defer db.Close()

	commentRepository := repository.NewCommentRepository(db)
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	db := connection.GetConnection()
	defer db.Close()

	commentRepository := repository.NewCommentRepository(db)
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}

func TestCustomerInsert(t *testing.T) {
	db := connection.GetConnection()
	defer db.Close()

	customerRepository := repository.NewCustomerRepository(db)
	ctx := context.Background()

	customer := entity.Customer{
		Id:   "C001",
		Name: "John Doe",
	}
	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
