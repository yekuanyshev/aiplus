package response

type City struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CityCreate struct {
	ID int64 `json:"id"`
}
