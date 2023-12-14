package repository

import (
	"context"
	"pinnacle-play/domain/model"
)

type GroupRepository interface {
	Save(ctx context.Context, name string) (*model.Group, error)
}