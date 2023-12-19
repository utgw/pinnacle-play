package output

import "pinnacle-play/domain/model"

type PostGroupOutput struct {
	Group     *model.Group
	Users     []*model.User
	Questions []*model.Question
}
