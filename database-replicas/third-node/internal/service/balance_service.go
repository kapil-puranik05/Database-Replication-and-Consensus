package service

import (
	"errors"
	"math"
	"third-node/internal/dto"
	"third-node/internal/model"
	"third-node/internal/repository"
)

func CreateBalance(balanceRequest dto.BalanceRequest) (model.Balance, error) {
	balance := model.Balance{
		Balance: balanceRequest.Balance,
	}
	return repository.Save(balance)
}

func UpdateBalance(balanceUpdateRequest dto.BalanceUpdateRequest) (model.Balance, error) {
	balance, err := repository.GetById(balanceUpdateRequest.UserId)
	if err != nil {
		return balance, errors.New("User Id not found")
	}
	if balanceUpdateRequest.Amount < 0 && balance.Balance < math.Abs(balanceUpdateRequest.Amount) {
		return balance, errors.New("Insufficient balance")
	}
	balance.Balance += balanceUpdateRequest.Amount
	return repository.Save(balance)
}

func GetBalance(id uint) (model.Balance, error) {
	return repository.GetById(id)
}

func GetAllBalances() ([]model.Balance, error) {
	return repository.GetAll()
}

func DeleteBalance(id uint) error {
	return repository.DeleteByID(id)
}
