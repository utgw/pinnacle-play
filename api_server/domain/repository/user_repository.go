package repository

import (
	"context"
	"pinnacle-play/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, name model.UserName, groupId model.GroupID) (*model.User, error)
}
