package service

import (
	"cryptodini-robo-manager/model"
	"fmt"
)

// ICryptodiniService interface
type ICryptodiniService interface {
	Adjust(uid string, desirePort *model.Portfolio)
	GetPort(uid string) *model.Portfolio
}

// CryptodiniService model
type CryptodiniService struct {
	AssetManagement IAssetManagerService
}

// NewCryptodiniService entity
func NewCryptodiniService() (ICryptodiniService, error) {
	return &CryptodiniService{}, nil
}

// Adjust adjust port via cryptodini
func (c *CryptodiniService) Adjust(uid string, desirePort *model.Portfolio) {
	fmt.Println("send http request to cryptodini to adjust port based on algorithm")
}

// GetPort get port value via cryptodini
func (c *CryptodiniService) GetPort(uid string) *model.Portfolio {
	fmt.Println("get currnet port value from binance")
	return &model.Portfolio{}
}
