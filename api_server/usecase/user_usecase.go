package usecase

import (
	"context"
	"errors"
	"pinnacle-play/domain/model"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase/input"
)

type User interface {
	CreateUser(ctx context.Context, in input.PostUserInput, groupId model.GroupID) ([]*model.User, error)
	ValidatePost(in input.PostUserInput) error
}
type userUsecase struct {
	userRepo repository.UserRepository
}

// Ensure userUsecase implements User interface at compile time
var _ User = (*userUsecase)(nil)

func NewUserUsecase(u repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, in input.PostUserInput, groupId model.GroupID) ([]*model.User, error) {
	var users []*model.User
	for _, userName := range in.UserNames {
		user, err := u.userRepo.Save(ctx, userName, groupId)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userUsecase) ValidatePost(in input.PostUserInput) error {
	if len(in.UserNames) == 0 {
		return errors.New("UserNames parameter is invalid")
	}
	for _, userName := range in.UserNames {
		if userName == "" {
			return errors.New("UserNames parameter is invalid")
		}
	}
	return nil
}
