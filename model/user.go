package model

// User model
type User struct {
	UserID       string
	Name         string
	WalletAdress string
	Ports        []Portfolio
}
