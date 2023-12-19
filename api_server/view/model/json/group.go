package json

// Group Group's Out Entity
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User User's Out Entity
type User struct {
	Name string `json:"user_name"`
}

// Question Question's Out Entity
type Question struct {
	Content string `json:"question_content"`
}

// PostGroupResult Out Entity For PostGroup Result
type PostGroupResult struct {
	Group     *Group      `json:"group"`
	Users     []*User     `json:"users"`
	Questions []*Question `json:"questions"`
}
