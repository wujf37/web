package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type OrderProduct struct {
	Id            uint   `json:"id"`                     // Order ID corresponds to Order struct
	ProductId     uint   `json:"product_id"`             // Product ID corresponds to Order struct
	ProductName   string `json:"product_name"`           // Product name corresponds to Product struct
	ImgPath       string `json:"img_path"`               // ImgPath corresponds to Product struct
	Num           int    `json:"num"`                    // Num corresponds to Order struct
	Price         string `json:"product_price"`          // Price corresponds to Product struct
	DiscountPrice string `json:"product_discount_price"` // DiscountPrice corresponds to Product struct
}

type PaginationOrderRequest struct {
	//Page        int    `json:"page"`
	//NumEachPage int    `json:"num_each_page"`
	UserID int `json:"user_id"`
	//Status      string `json:"status"`
}

type OrderDisplay struct {
	ID               uint           `json:"id"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	OrderID          uint64         `json:"order_id"`
	UserName         string         `json:"user_name"`
	UserId           int            `json:"user_id"`
	Address          string         `json:"address"`
	UserPhone        string         `json:"user_phone"`
	Status           string         `json:"status"`
	CanteenID        int            `json:"canteen_id"`
	DeliverID        int            `json:"deliver_id"`
	DeliverName      string         `json:"deliver_name"`
	DeliverPhone     string         `json:"deliver_phone"`
	OrderProductList []OrderProduct `json:"order_productlist"`
}

func GetOrder(c *gin.Context) {
	var pagination PaginationOrderRequest
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(400, gin.H{"msg": "数据格式错误", "status": 201})
		return
	}

	var orders []DBstruct.Order
	// if pagination.Status != "" {
	// 	DBstruct.DB.Where("user_id = ? AND status = ?", pagination.UserID, pagination.Status).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
	// } else {
	//	DBstruct.DB.Where("user_id = ?", pagination.UserID).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
	// }
	DBstruct.DB.Where("user_id = ?", pagination.UserID).Order("updated_at").Find(&orders)
	var orderDisplays []OrderDisplay
	var orderDisplay OrderDisplay
	for _, order := range orders {
		if orderDisplay.OrderID != order.OrderID {
			if orderDisplay.OrderID != 0 {
				orderDisplays = append(orderDisplays, orderDisplay)
			}
			orderDisplay = OrderDisplay{
				ID:           order.ID,
				CreatedAt:    order.CreatedAt.String(),
				UpdatedAt:    order.UpdatedAt.String(),
				OrderID:      order.OrderID,
				UserName:     order.UserName,
				UserId:       order.UserId,
				Status:       order.Status,
				CanteenID:    order.CanteenID,
				DeliverID:    order.DeliverID,
				DeliverName:  order.DeliverName,
				DeliverPhone: order.DeliverPhone,
				Address:      order.Address,
				UserPhone:    order.UserPhone,
			}
		}

		var product DBstruct.Product
		DBstruct.DB.Where("id = ?", order.ProductId).First(&product)

		orderProduct := OrderProduct{
			Id:            order.ID,
			ProductId:     product.ID,
			ProductName:   product.Name,
			ImgPath:       product.ImgPath,
			Num:           order.Num,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
		}

		// Add order product to order's product list
		orderDisplay.OrderProductList = append(orderDisplay.OrderProductList, orderProduct)
	}
	if orderDisplay.OrderID != 0 {
		orderDisplays = append(orderDisplays, orderDisplay)
	}
	// var totalCount int
	// if pagination.Status != "" {
	// 	DBstruct.DB.Model(&DBstruct.Order{}).Where("user_id = ? AND status = ?", pagination.UserID, pagination.Status).Count(&totalCount)
	// } else {
	// 	DBstruct.DB.Model(&DBstruct.Order{}).Where("user_id = ?", pagination.UserID).Count(&totalCount)
	// }

	c.JSON(200, gin.H{
		"msg":    "OK",
		"status": 200,
		"data": gin.H{
			// "count":     totalCount,
			"orderlist": orderDisplays,
		},
	})
}
