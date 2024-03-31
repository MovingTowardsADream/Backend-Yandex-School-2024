package entity

type Courier struct {
	Id        int      `json:"id"`
	Type      string   `json:"type" binding:"required"`
	Districts []int    `json:"districts" binding:"required"`
	Schedule  []string `json:"schedule" binding:"required"`
}

type Couriers struct {
	Couriers []Courier `json:"couriers"`
}
