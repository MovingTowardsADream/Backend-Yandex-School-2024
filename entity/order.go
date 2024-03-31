package entity

type Order struct {
	Id             int      `json:"id"`
	Weight         int      `json:"weight"`
	District       int      `json:"district"`
	ConvenientTime []string `json:"convenient_time"`
}
