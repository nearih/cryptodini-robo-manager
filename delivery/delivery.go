package delivery

import (
	"cryptodini-robo-manager/service"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	CryptodiniService    service.ICryptodiniService
	IAssetManagerService service.IAssetManagerService
}

type assetManagementInput struct {
	Uid    string
	Amount float64
}

// NewDelivery create new delivery entity
func NewDelivery(asset service.IAssetManagerService) (*Handler, error) {
	return &Handler{
		IAssetManagerService: asset,
	}, nil
}

// Test for connection testing
func (h *Handler) Test() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, "success")
	}
}

// Deposit asset to exchange
func (h *Handler) Deposit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req assetManagementInput
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		buyOrder := h.IAssetManagerService.Deposit(req.Uid, req.Amount)

		return c.JSON(http.StatusOK, buyOrder)
	}
}

// Withdraw withdraw asset from exchange
func (h *Handler) Withdraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req assetManagementInput
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		sellOrder := h.IAssetManagerService.Withdraw(req.Uid, req.Amount)

		return c.JSON(http.StatusOK, sellOrder)
	}
}

// GetPort return current port value
func (h *Handler) GetPort() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req assetManagementInput
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		port := h.IAssetManagerService.GetPort(req.Uid)

		return c.JSON(http.StatusOK, port)
	}
}

// // Adjust rebalance port
// func (h *Handler) Adjust() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		h.CryptodiniService.Adjust("mock uid", &model.Portfolio{})

// 		return c.JSON(http.StatusOK, "adjusted")
// 	}
// }

// // GetPort get port value
// func (h *Handler) GetPort() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		p := h.CryptodiniService.GetPort("mock uid")

// 		return c.JSON(http.StatusOK, p.ID)
// 	}
// }
