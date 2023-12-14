package model

import "time"

type User struct {
	id            UserID
	name          UserName
	groupId 			GroupID
	createdAt       time.Time
}

type UserName  string

type UserID int

type Users  []User


func NewUser(id UserID, name UserName, groupId GroupID, createdAt time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		groupId:   groupId,
		createdAt: createdAt,
	}
}

func (u User) ID() UserID {
	return u.id
}


func (u User) Name() UserName {
	return u.name
}

func (u User) GroupId() GroupID {
	return u.groupId
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}
