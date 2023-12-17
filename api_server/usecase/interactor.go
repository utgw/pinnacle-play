package usecase

import (
	"context"
	"pinnacle-play/usecase/input"
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

// PostGroup save memo and tags
func (i Interactor) PostGroup(ctx context.Context, in input.PostGroupInput) {
	err := i.group.ValidatePost(ctx, in)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	group, err := i.group.CreateGroup(ctx, in)
	if err != nil {
		i.pre.ViewError(ctx, err)
		return
	}

	i.pre.ViewGroup(ctx, group)
}
