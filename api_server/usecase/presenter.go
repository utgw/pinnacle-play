package usecase

import (
	"context"
	"pinnacle-play/domain/model"
)

// Presenter presenter interface
type Presenter interface {
	ViewError(ctx context.Context, err error)
	ViewGroup(ctx context.Context, mg *model.Group)
}
