package output

import "pinnacle-play/domain/model"

type PostGroupOutput struct {
	Group     *model.Group
	Users     *model.Users
	Questions *model.Questions
}
