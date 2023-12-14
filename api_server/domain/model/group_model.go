package model

import (
	"math/rand"
	"time"
)

type GroupID string

type GroupName string

type Group struct {
	id        GroupID
	name      GroupName
	createdAt time.Time
}

type Groups []Group

func NewGroup(name GroupName, createdAt time.Time) *Group {
	return &Group{
		id:        GroupID(generateRandomID(groupIdLength)),
		name:      name,
		createdAt: createdAt,
	}
}

func (g Group) ID() GroupID {
	return g.id
}

func (g Group) Name() GroupName {
	return g.name
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

const groupIdLength int = 16
// generateRandomIDは指定された長さのランダムな英数字の文字列を生成します。
func generateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}