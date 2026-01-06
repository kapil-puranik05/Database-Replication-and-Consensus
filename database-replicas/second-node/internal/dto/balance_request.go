package dto

type BalanceRequest struct {
	Balance float64 `json:"balance" binding:"required,gte=0"`
}
