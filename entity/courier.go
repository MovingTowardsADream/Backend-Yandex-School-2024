package entity

type Courier struct {
	Id        int      `json:"id"`
	Type      string   `json:"type"`
	Districts []int    `json:"districts"`
	Schedule  []string `json:"schedule"`
}
