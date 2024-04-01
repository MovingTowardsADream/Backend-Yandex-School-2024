package entity

type History struct {
	Id        int    `json:"id"`
	CourierId int    `json:"courier_id" binding:"required"`
	OrderId   int    `json:"order_id" binding:"required"`
	Time      string `json:"time" binding:"required"`
	Date      string `json:"date" binding:"required"`
}

type Histories struct {
	Histories []History `json:"histories"`
}
