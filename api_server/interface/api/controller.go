package api

import (
	"net/http"
	"pinnacle-play/domain/model"
	"pinnacle-play/infra/logger"
	"pinnacle-play/usecase/input"
	"pinnacle-play/usecase/interactor"
)

// NewAPI Get API instance
func NewAPI(it interactor.Interactor, log logger.Logger) API {
	return controller{it, log}
}

// API api instance
type API interface {
	PostGroup(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	it  interactor.Interactor
	log logger.Logger
}

var _ API = (*controller)(nil)

func (c controller) PostGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = addResponseWriter(ctx, w)

	userNames := c.convertToUserNames(r.URL.Query()["userNames"])
	questionContents := c.convertToQuestionContents(r.URL.Query()["questionContents"])

	in := &input.CreateGroupInput{
		GroupInput:    input.PostGroupInput{GroupName: model.GroupName(r.URL.Query().Get("groupName"))},
		UserInput:     input.PostUserInput{UserNames: userNames},
		QuestionInput: input.PostQuestonInput{QuestionContent: questionContents},
	}
	c.it.CreateGroup(ctx, *in)
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
