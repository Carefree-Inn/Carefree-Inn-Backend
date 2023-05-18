package handler

import (
	"user-post/internal/repository"
)

type UserPostService struct {
	userPostDao repository.UserPostRepository
}

func NewUserPostService(database *repository.Database) *UserPostService {
	return &UserPostService{
		userPostDao: repository.NewUserPostRepository(database),
	}
}
