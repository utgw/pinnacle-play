package usecase

import (
	"context"
	"errors"

	"pinnacle-play/domain/model"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase/input"
)

type Group interface {
	CreateGroup(ctx context.Context, ipt input.PostGroupInput) (*model.Group, error)
	ValidatePost(ipt input.PostGroupInput) error
}

type groupUsecase struct {
	groupRepo repository.GroupRepository
}

var _ Group = (*groupUsecase)(nil)

func NewGroupUsecase(g repository.GroupRepository) *groupUsecase {
	return &groupUsecase{
		groupRepo: g,
	}
}

func (g *groupUsecase) ValidatePost(in input.PostGroupInput) error {
	if in.GroupName == "" {
		return errors.New("GroupName parameter is invalid")
	}
	return nil
}

func (u *groupUsecase) CreateGroup(ctx context.Context, in input.PostGroupInput) (*model.Group, error) {
	group, err := u.groupRepo.Save(ctx, in.GroupName)
	if err != nil {
		return nil, err
	}

	return group, nil
}
