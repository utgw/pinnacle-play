package model

import "time"


type Question struct {
	id        QuestionID
	content     QuestionContent
	groupId GroupID
	createdAt time.Time
}

type QuestionContent string

type QuestionID int

type Questions []Question

func NewQuestion(id QuestionID, content QuestionContent, groupId GroupID, createdAt time.Time) *Question {
	return &Question{
		id:        id,
		content:     content,
		groupId:    groupId,
		createdAt: createdAt,
	}
}

func (q Question) ID() QuestionID {
	return q.id
}

func (q Question) Content() QuestionContent {
	return q.content
}

func (q Question) GroupID() GroupID {
	return q.groupId
}

func (q Question) CreatedAt() time.Time {
	return q.createdAt
}
