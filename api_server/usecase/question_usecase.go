package usecase

import (
	"context"
	"errors"
	"pinnacle-play/domain/model"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase/input"
)

type Question interface {
	CreateQuestion(ctx context.Context, in input.PostQuestonInput, groupId model.GroupID) ([]*model.Question, error)
	ValidatePost(in input.PostQuestonInput) error
}

type questionUsecase struct {
	questionRepo repository.QuestionRepository
}

// Ensure questionUsecase implements Question interface at compile time
var _ Question = (*questionUsecase)(nil)

func NewQuestionUsecase(q repository.QuestionRepository) *questionUsecase {
	return &questionUsecase{
		questionRepo: q,
	}
}

func (u *questionUsecase) CreateQuestion(ctx context.Context, in input.PostQuestonInput, groupId model.GroupID) ([]*model.Question, error) {
	var questions []*model.Question
	questionContents := in.QuestionContent
	for _, questionContent := range questionContents {
		question, err := u.questionRepo.Save(ctx, questionContent, groupId)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (u *questionUsecase) ValidatePost(in input.PostQuestonInput) error {
	for _, questionContent := range in.QuestionContent {
		if questionContent == "" {
			return errors.New("QuestionContents parameter is invalid")
		}
	}

	return nil
}
