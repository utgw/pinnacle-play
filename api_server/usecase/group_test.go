package usecase

import (
	"errors"
	mock_repository "pinnacle-play/domain/repository/mocks"
	"pinnacle-play/usecase/input"
	"testing"
)

func TestGroupUsecase_ValidatePost(t *testing.T) {
	userRepo := new(mock_repository.MockUserRepository)
	groupRepo := new(mock_repository.MockGroupRepository)
	questionRepo := new(mock_repository.MockQuestionRepository)
	transactionRepo := new(mock_repository.MockTransactionRepository)
	u := NewGroupUsecase(userRepo, groupRepo, questionRepo, transactionRepo)
	// テストケースを定義します
	tests := []struct {
		name  string
		input input.PostGroupInput
		want  error
	}{
		{
			name: "全てのバリデートをパスするとき",
			input: input.PostGroupInput{
				GroupName:        "group1",
				UserNames:        []string{"user1", "user2"},
				QuestionContents: []string{"question1", "question2"},
			},
			want: nil,
		},
		{
			name: "GroupNameが空のとき",
			input: input.PostGroupInput{
				GroupName:        "",
				UserNames:        []string{"user1", "user2"},
				QuestionContents: []string{"question1", "question2"},
			},
			want: errors.New("GroupName parameter is invalid"),
		},
		{
			name: "UserNamesの要素が0のとき",
			input: input.PostGroupInput{
				GroupName:        "group1",
				UserNames:        []string{},
				QuestionContents: []string{"question1", "question2"},
			},
			want: errors.New("UserNames parameter is invalid"),
		},
		{
			name: "usernameの要素が0ではないがその中の要素が空のとき",
			input: input.PostGroupInput{
				GroupName:        "group1",
				UserNames:        []string{"", "user2"},
				QuestionContents: []string{"question1", "question2"},
			},
			want: errors.New("UserNames parameter is invalid"),
		},
		{
			name: "questionContentsの要素が0ではないがその中の要素が空のとき",
			input: input.PostGroupInput{
				GroupName:        "group1",
				UserNames:        []string{"user1", "user2"},
				QuestionContents: []string{"", "question2"},
			},
			want: errors.New("QuestionContents parameter is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := u.ValidatePost(tt.input)
			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("got %v, want %v", err, tt.want)
				}
			} else if tt.want != nil {
				t.Errorf("got nil, want %v", tt.want)
			}
		})
	}
}
