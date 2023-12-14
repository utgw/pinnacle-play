package repository

import (
	"context"
	"pinnacle-play/domain/model"
)

type QuestionRepository interface {
	Save(ctx context.Context, content string, groupId model.GroupID) error
}