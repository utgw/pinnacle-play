package interactor

import (
	"context"
	"pinnacle-play/domain/repository"
	"pinnacle-play/usecase"
	"pinnacle-play/usecase/input"
	"pinnacle-play/usecase/output"
)

type Interactor interface {
	CreateGroup(ctx context.Context, ipt input.CreateGroupInput)
}

type createGroupInteractor struct {
	presenter       usecase.Presenter
	userUsecase     usecase.User
	questionUsecase usecase.Question
	groupUsecase    usecase.Group
	transactionRepo repository.TransactionRepository
}

var _ Interactor = (*createGroupInteractor)(nil)

func NewCreateGroupInteractor(presenter usecase.Presenter, userUsecase usecase.User, questionUsecase usecase.Question, groupUsecase usecase.Group, transactionRepo repository.TransactionRepository) *createGroupInteractor {
	return &createGroupInteractor{
		presenter:       presenter,
		userUsecase:     userUsecase,
		questionUsecase: questionUsecase,
		groupUsecase:    groupUsecase,
		transactionRepo: transactionRepo,
	}
}

func (i *createGroupInteractor) CreateGroup(ctx context.Context, input input.CreateGroupInput) {
	// トランザクション開始
	txCtx, transactErr := i.transactionRepo.Begin(ctx)
	if transactErr != nil {
		i.presenter.ViewError(ctx, transactErr)
		return
	}

	err := i.groupUsecase.ValidatePost(input.GroupInput)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	group, err := i.groupUsecase.CreateGroup(ctx, input.GroupInput)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	err = i.userUsecase.ValidatePost(input.UserInput)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	users, err := i.userUsecase.CreateUser(ctx, input.UserInput, group.ID())
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	err = i.questionUsecase.ValidatePost(input.QuestionInput)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	questions, err := i.questionUsecase.CreateQuestion(ctx, input.QuestionInput, group.ID())
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	// 処理が成功したらコミット、失敗したらロールバック
	if err != nil {
		_, rbErr := i.transactionRepo.Rollback(txCtx)
		if rbErr != nil {
			i.presenter.ViewError(ctx, rbErr)
		}
		i.presenter.ViewError(ctx, err)
	} else {
		_, cErr := i.transactionRepo.Commit(txCtx)
		if cErr != nil {
			i.presenter.ViewError(ctx, cErr)
		}
	}

	output := output.PostGroupOutput{
		Group:     group,
		Users:     users,
		Questions: questions,
	}

	i.presenter.ViewPostGroupResult(ctx, output)
}
