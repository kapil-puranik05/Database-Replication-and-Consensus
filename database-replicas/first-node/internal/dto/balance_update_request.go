package dto

type BalanceUpdateRequest struct {
	UserId uint    `json:"userId"`
	Amount float64 `json:"amount" binding:"required,gte=0"`
}
