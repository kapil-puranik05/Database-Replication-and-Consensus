package repository

import (
	"first-node/internal/database"
	"first-node/internal/model"
)

func Save(balance model.Balance) (model.Balance, error) {
	err := database.DB.Create(&balance).Error
	return balance, err
}

func GetById(id uint) (model.Balance, error) {
	var balance model.Balance
	err := database.DB.First(&balance, id).Error
	return balance, err
}

func GetAll() ([]model.Balance, error) {
	var balances []model.Balance
	err := database.DB.Find(&balances).Error
	return balances, err
}

func DeleteByID(id uint) error {
	return database.DB.Delete(&model.Balance{}, id).Error
}
