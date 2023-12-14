package usecase

import (
	"context"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase/input"
)

type Group interface {
	CreateGroupWithUsersAndQuestions(ctx context.Context, ipt input.PostGroupWithUsersAndQuestionsInput) error
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

func (u *GroupUsecase) CreateGroupWithUsersAndQuestions(ctx context.Context, in input.PostGroupWithUsersAndQuestionsInput) error {
	tx := u.txRepo
	// トランザクションを開始
	ctx, err := tx.Begin(ctx)
	if err != nil {
		return err
	}

	// groupを保存
	group, err := u.groupRepo.Save(ctx, in.GroupName)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	// 各userとquestionにgroupのIDを設定し、保存
	userNames := in.UserNames
	for _, userName := range userNames {
		if err := u.userRepo.Save(ctx, userName, group.ID()); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	questionContents := in.QuestionContents
	for _, questionContent := range questionContents {
		if err := u.questionRepo.Save(ctx, questionContent, group.ID()); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	// トランザクションをコミット
	if _,err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}