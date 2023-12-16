package usecase

import (
	"context"
	"pinnacle-play/domain/model"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase/input"
)

type Group interface {
	CreateGroup(ctx context.Context, ipt input.PostGroupInput) (*model.Group, error)
	ValidatePost(ctx context.Context, ipt input.PostGroupInput) error
}

type GroupUsecase struct {
	userRepo     repository.UserRepository
	groupRepo    repository.GroupRepository
	questionRepo repository.QuestionRepository
	txRepo       repository.TransactionRepository
}

func NewGroupUsecase(u repository.UserRepository, g repository.GroupRepository, q repository.QuestionRepository,t repository.TransactionRepository) *GroupUsecase {
	return &GroupUsecase{
		userRepo:     u,
		groupRepo:    g,
		questionRepo: q,
		txRepo:       t,
	}
}

func (u *GroupUsecase) CreateGroup(ctx context.Context, in input.PostGroupInput) (*model.Group, error) {
	tx := u.txRepo
	// トランザクションを開始
	ctx, err := tx.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// groupを保存
	group, err := u.groupRepo.Save(ctx, in.GroupName)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	// 各userとquestionにgroupのIDを設定し、保存
	userNames := in.UserNames
	for _, userName := range userNames {
		if err := u.userRepo.Save(ctx, userName, group.ID()); err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
	}

	questionContents := in.QuestionContents
	for _, questionContent := range questionContents {
		if err := u.questionRepo.Save(ctx, questionContent, group.ID()); err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
	}

	// トランザクションをコミット
	if _,err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return group, nil
}