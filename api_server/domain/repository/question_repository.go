package repository

import (
	"context"
)

type QuestionRepository interface {
	Save(ctx context.Context, content string) error
}