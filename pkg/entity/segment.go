package entity

type Segment struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Metadata    []string `json:"metadata"`
	UserList    []int    `json:"user_list"`
}
