package entity

type CourierRating struct {
	Earnings int `json:"earnings" binding:"required"`
	Rating   int `json:"rating" binding:"required"`
}

type Period struct {
	StartDate string
	EndDate   string
}

type CourierMeta struct {
	Count int
	Sum   int
}
