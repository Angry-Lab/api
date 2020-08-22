package city

type City struct {
	ID        int     `json:"id"`
	Postcode  int     `json:"postcode"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
