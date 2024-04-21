package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Item struct {
	InstanceID   string `json:"instance_id"`
	Hostname     string `json:"hostname"`
	InstanceType string `json:"instance_type"`
	Platform     string `json:"platform"`
}

var (
	mutex sync.Mutex
	items []Item
)

func main() {
	r := gin.Default()

	// 添加CORS中间件
	r.Use(cors.Default())

	// 添加数据
	r.POST("/api/v1/cmdb", addItem)

	// 获取所有数据
	r.GET("/api/v1/cmdb", getAllItems)

	// 根据InstanceID获取单个数据
	r.GET("/api/v1/cmdb/:instance_id", getItemByInstanceID)

	// 更新数据
	r.PUT("/api/v1/cmdb/:instance_id", updateItem)

	// 删除数据
	r.DELETE("/api/v1/cmdb/:instance_id", deleteItem)

	r.Run(":8080")
}

func addItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	items = append(items, newItem)

	fmt.Printf("Added new item: %+v\n", newItem)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": newItem, "message": "create ok"})
}

func getAllItems(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": items, "message": ""})
}

func getItemByInstanceID(c *gin.Context) {
	instanceID := c.Param("instance_id")

	mutex.Lock()
	defer mutex.Unlock()

	for _, item := range items {
		if item.InstanceID == instanceID {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": item, "message": ""})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func updateItem(c *gin.Context) {
	instanceID := c.Param("instance_id")

	mutex.Lock()
	defer mutex.Unlock()

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.InstanceID == instanceID {
			items[i] = updatedItem
			c.JSON(http.StatusOK, items[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteItem(c *gin.Context) {
	instanceID := c.Param("instance_id")

	mutex.Lock()
	defer mutex.Unlock()

	for i, item := range items {
		if item.InstanceID == instanceID {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": "删除成功"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

