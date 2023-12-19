package repository

import (
	"context"
	"pinnacle-play/domain/model"
)

type QuestionRepository interface {
	Save(ctx context.Context, content model.QuestionContent, groupId model.GroupID) (*model.Question, error)
}
