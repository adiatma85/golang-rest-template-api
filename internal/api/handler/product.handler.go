package handler

import (
	"net/http"
	"strconv"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/repository"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/validator"
	"github.com/adiatma85/go-tutorial-gorm/pkg/helpers"
	"github.com/adiatma85/go-tutorial-gorm/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

// Func to Create User, similar to #Register
func CreateProduct(c *gin.Context) {
	var createProductRequest validator.CreateProductRequest
	err := c.ShouldBind(&createProductRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to create new product", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	productRepo := repository.GetProductRepository()
	productModel := &models.Product{}

	// smapping the struct
	smapping.FillStruct(productModel, smapping.MapFields(&createProductRequest))

	if newProduct, err := productRepo.Create(*productModel); err != nil {
		response := response.BuildFailedResponse("failed to create new product", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		response := response.BuildSuccessResponse("success to create new product", newProduct)
		c.JSON(http.StatusOK, response)
		return
	}
}

// Func to GetAll User without in server pagination
func GetAllProduct(c *gin.Context) {
	productRepo := repository.GetProductRepository()

	products, err := productRepo.GetAll()
	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", products)
	c.JSON(http.StatusOK, response)
}

// Func to GetSpecific Product
func GetSpecificProduct(c *gin.Context) {
	productRepo := repository.GetProductRepository()

	product, err := productRepo.GetById(c.Param("productId"))

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", product)
	c.JSON(http.StatusOK, response)
}

// Func to Query Product with pagination
func QueryProducts(c *gin.Context) {
	pagination := helpers.Pagination{}
	productRepo := repository.GetProductRepository()
	queryPageLimit, isPageLimitExist := c.GetQuery("limit")
	queryPage, isPageQueryExist := c.GetQuery("page")

	if isPageQueryExist {
		pagination.Page, _ = strconv.Atoi(queryPage)
	} else {
		pagination.Page = 1
	}

	if isPageLimitExist {
		pagination.Limit, _ = strconv.Atoi(queryPageLimit)
	} else {
		pagination.Limit = 10
	}

	products, err := productRepo.Query(&models.Product{}, pagination)

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", products)
	c.JSON(http.StatusOK, response)
}

// Func to Update Product,
func UpdateSpecificProduct(c *gin.Context) {
	var updateRequest validator.UpdateProductRequest
	err := c.ShouldBind(&updateRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a product", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	updateModel := &models.Product{}

	// smapping the update request to models
	updateModel.ID, _ = strconv.ParseUint(c.Param("productId"), 10, 64)
	smapping.FillStruct(updateModel, smapping.MapFields(&updateRequest))

	productRepo := repository.GetProductRepository()
	err = productRepo.Update(updateModel)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a product", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Specific Product
func DeleteSpecificProduct(c *gin.Context) {
	deleteModel := &models.Product{}
	deleteModel.ID, _ = strconv.ParseUint(c.Param("productId"), 10, 64)

	productRepo := repository.GetProductRepository()

	err := productRepo.Delete(deleteModel)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete a product", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Products with array ids
func DeleteProductsWithIds(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "ok",
		"message": "need revision for delete product",
	})
}
