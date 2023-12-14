package repository

import (
	"context"
	"pinnacle-play/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, name string, groupId model.GroupID) error
}