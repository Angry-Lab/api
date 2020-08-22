package segment

type Segment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Condition   string `json:"condition"`

	Metadata Metadata `json:"metadata"`
	//Metrics  Metrics  `json:"metrics"`
	//UserList []int    `json:"user_list"`
}

type Metadata struct {
	TotalCost float64 `json:"total_cost"`
	TotalNP   float64 `json:"total_np"`
	Count     int     `json:"count"`
	Users     int     `json:"users"`
	Distance  float64 `json:"distance"`
}

type Metrics struct {
	Activity []float64 `json:"activity"`
	Solvency []float64 `json:"solvency"`
	NP       []float64 `json:"np"`
	Weight   []float64 `json:"weight"`
}
