package usecase

import (
	"context"
	"pinnacle-play/usecase/output"
)

// Presenter presenter interface
type Presenter interface {
	ViewError(ctx context.Context, err error)
	ViewPostGroupResult(ctx context.Context, output output.PostGroupOutput)
}
