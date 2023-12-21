package api

import (
	"net/http"
	"pinnacle-play/domain/model"
	"pinnacle-play/infra/logger"
	"pinnacle-play/usecase"
	"pinnacle-play/usecase/input"
)

// NewAPI Get API instance
func NewAPI(it usecase.Interactor, log logger.Logger) API {
	return controller{it, log}
}

// API api instance
type API interface {
	PostGroup(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	it  usecase.Interactor
	log logger.Logger
}

var _ API = (*controller)(nil)

func (c controller) PostGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = addResponseWriter(ctx, w)

	userNames := c.convertToUserNames(r.URL.Query()["userNames"])
	questionContents := c.convertToQuestionContents(r.URL.Query()["questionContents"])

	in := &input.PostGroupInput{
		GroupName:        model.GroupName(r.URL.Query().Get("groupName")),
		UserNames:        userNames,
		QuestionContents: questionContents,
	}
	c.it.PostGroup(ctx, *in)
}

func (c controller) convertToUserNames(names []string) []model.UserName {
	converted := make([]model.UserName, len(names))
	for i, name := range names {
		converted[i] = model.UserName(name)
	}
	return converted
}

func (c controller) convertToQuestionContents(contents []string) []model.QuestionContent {
	converted := make([]model.QuestionContent, len(contents))
	for i, content := range contents {
		converted[i] = model.QuestionContent(content)
	}
	return converted
}
