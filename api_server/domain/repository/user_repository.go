package repository

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, name string) error
}