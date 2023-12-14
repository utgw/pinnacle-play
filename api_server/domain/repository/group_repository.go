package repository

import (
	"context"
)

type GroupRepository interface {
	Save(ctx context.Context, name string) error
}