package controller

import (
	"go-first/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order model.Order
		res   gin.H
	)

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := idb.DB.Create(&order).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := range order.Items {
		order.Items[i].OrderID = order.OrderId
		if err := idb.DB.Create(&order.Items[i]).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	res = gin.H{
		"message": "Data pesanan berhasil ditambahkan",
	}

	c.JSON(http.StatusCreated, res)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []model.Order
		items  []model.Item
		result gin.H
	)

	if err := idb.DB.Find(&orders).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := range orders {
		if err := idb.DB.Where("order_id = ?", orders[i].OrderId).Find(&items).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		orders[i].Items = items
	}

	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrder(c *gin.Context) {
	var (
		order  model.Order
		items  []model.Item
		result gin.H
	)

	id := c.Param("id")

	if err := idb.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := idb.DB.Where("order_id = ?", order.OrderId).Find(&items).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	order.Items = items

	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var (
		order  model.Order
		result gin.H
	)

	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"message": "Data Tidak Ditemukan",
		}
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := idb.DB.Save(&order).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, newItem := range order.Items {
		var existingItem model.Item
		if err := idb.DB.Where("item_id = ?", newItem.ItemId).First(&existingItem).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		existingItem.ItemCode = newItem.ItemCode
		existingItem.Description = newItem.Description
		existingItem.Quantity = newItem.Quantity

		if err := idb.DB.Save(&existingItem).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	result = gin.H{
		"result": "Data pesanan berhasil diperbarui",
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  model.Order
		items  []model.Item
		result gin.H
	)

	id := c.Param("id")

	if err := idb.DB.First(&order, id).Error; err != nil {
		result = gin.H{
			"result": "Data Tidak Ditemukan",
		}
	}

	if err := idb.DB.Where("order_id = ?", id).Delete(&items).Error; err != nil {
		result = gin.H{
			"result": "Data Tidak Ditemukan",
		}
	}

	if err := idb.DB.Delete(&order).Error; err != nil {
		result = gin.H{
			"result": "Data Gagal Dihapus",
		}
	} else {
		result = gin.H{
			"result": "Data Berhasil Dihapus",
		}
	}

	c.JSON(http.StatusAccepted, result)
}
