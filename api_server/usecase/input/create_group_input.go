package input

import "pinnacle-play/domain/model"

type CreateGroupInput struct {
	GroupInput    PostGroupInput
	UserInput     PostUserInput
	QuestionInput PostQuestonInput
}

// 以下の置き場所ここか？
type PostGroupInput struct {
	GroupName model.GroupName
}

type PostUserInput struct {
	UserNames []model.UserName
}

type PostQuestonInput struct {
	QuestionContent []model.QuestionContent
}
