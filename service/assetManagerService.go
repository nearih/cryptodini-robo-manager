package service

import (
	"cryptodini-robo-manager/model"
	"fmt"
)

// IAssetManagerService interface
type IAssetManagerService interface {
	Deposit(uid string, amount float64) model.BuyOrder
	Withdraw(uid string, amount float64) model.SellOrder
	GetPort(uid string) model.Portfolio
}

// AssetManagerService model
type AssetManagerService struct {
	Cryptodini   ICryptodiniService
	AssetManager IAssetManagerService
}

// NewAssetManagerService entity
func NewAssetManagerService(crypSer ICryptodiniService) (IAssetManagerService, error) {
	return &AssetManagerService{
		Cryptodini: crypSer,
	}, nil
}

// Deposit allow asset manager to deposit fund
func (a *AssetManagerService) Deposit(uid string, amount float64) model.BuyOrder {
	fmt.Println("buy asset from binance")
	egOrder := make(map[string]float64)
	egOrder["btc"] = 50
	egOrder["xrp"] = 100

	// run first time it will create port for current user so it can auto adjust port
	go a.Cryptodini.Adjust(uid, &model.Portfolio{})

	return model.BuyOrder{
		ID:    "mock",
		Order: egOrder,
	}
}

// Withdraw allow asset manager to withdraw fund
func (a *AssetManagerService) Withdraw(uid string, amount float64) model.SellOrder {
	fmt.Println("sell asset to binance")
	egOrder := make(map[string]float64)
	egOrder["btc"] = 50
	egOrder["xrp"] = 100
	return model.SellOrder{
		ID:    "mock",
		Order: egOrder,
	}
}

// GetPort allow asset manager to get port value
func (a *AssetManagerService) GetPort(uid string) model.Portfolio {
	fmt.Println("get port value from binance")
	return model.Portfolio{}
}
