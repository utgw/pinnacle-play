package input

import "pinnacle-play/domain/model"

type PostGroupInput struct {
	GroupName        model.GroupName
	UserNames        []model.UserName
	QuestionContents []model.QuestionContent
}
