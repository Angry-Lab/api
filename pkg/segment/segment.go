package segment

type Segment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Condition   string `json:"condition"`

	Metadata []string `json:"metadata"`
	UserList []int    `json:"user_list"`
}
