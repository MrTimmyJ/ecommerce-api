// controllers/products.go

package controllers

import (
	"net/http"
	"time"

	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	PLU         int    `json:"plu" bindig:"required"`
}

// Get all products
func GetProducts(c *gin.Context) {
	prd := new(models.ProductRequestData)
	if err := models.DB.Find(&prd.Products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	prd.Retrieved = time.Now()
	c.JSON(http.StatusOK, prd)
}

func SearchProducts(c *gin.Context) {
	prd := new(models.ProductRequestData)
	query := c.Query("query")

	//Magic gorm search
	if err := models.DB.Where("id LIKE ? OR name LIKE ? OR type LIKE ? OR description LIKE ? OR image LIKE ? OR plu LIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&prd.Products).Error; err != nil {
		//Database error, return
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//Query completed succesfully, return json
	prd.Retrieved = time.Now()
	c.JSON(http.StatusOK, prd)
}

func CreateProduct(c *gin.Context) {
	// Validate input
	input := new(CreateProductInput)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create product
	product := models.Product{Name: input.Name, Type: input.Type, Description: input.Description,
		Image: input.Image, Quantity: input.Quantity, Price: input.Price,
		PLU: input.PLU}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
