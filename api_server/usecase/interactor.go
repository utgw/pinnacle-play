package usecase

import (
	"context"
	"pinnacle-play/usecase/input"
	"pinnacle-play/usecase/output"
)

// NewInteractor new Interactor
func NewInteractor(
	pre Presenter,
	group Group,
) Interactor {
	return Interactor{pre, group}
}

// Interactor usecase interactor
type Interactor struct {
	pre   Presenter
	group Group
}

// PostGroup 入力値の検証を行った後、グループを作成してpostします
func (i Interactor) PostGroup(ctx context.Context, in input.PostGroupInput) {
	err := i.group.ValidatePost(in)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	group, users, questions, err := i.group.CreateGroup(ctx, in)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	output := output.PostGroupOutput{
		Group:     group,
		Users:     users,
		Questions: questions,
	}

	i.pre.ViewPostGroupResult(ctx, output)
}
