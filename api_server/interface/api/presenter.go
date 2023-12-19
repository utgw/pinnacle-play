package api

import (
	"context"
	"encoding/json"
	"net/http"
	apperror "pinnacle-play/infra/error"
	"pinnacle-play/infra/logger"
	"pinnacle-play/usecase"
	"pinnacle-play/usecase/output"

	"pinnacle-play/view/render"
)

// NewPresenter new presenter
func NewPresenter(render render.JSONRender, log logger.Logger, errm apperror.ErrorManager) usecase.Presenter {
	return presenter{render, log, errm}
}

type presenter struct {
	render render.JSONRender
	log    logger.Logger
	errm   apperror.ErrorManager
}

func (p presenter) ViewPostGroupResult(ctx context.Context, output output.PostGroupOutput) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	p.JSON(ctx, w, p.render.ConvertPostGroupResult(output))
}

func (p presenter) ViewError(ctx context.Context, err error) {
	defer deleteResponseWriter(ctx)
	w := getResponseWriter(ctx)

	p.log.Errorf("API: %s\n", p.errm.LogMessage(err))

	p.JSON(ctx, w, p.render.ConvertError(err, p.errm.Code(err)))
}

// JSON render json format
func (p presenter) JSON(ctx context.Context, w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
