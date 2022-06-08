package handler

import (
	"net/http"
	"strconv"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/validator"
	"github.com/adiatma85/golang-rest-template-api/pkg/helpers"
	"github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

// Local variable
var (
	productHandler *ProductHandler
)

// Func to implement contract of productHandler
type ProductHandler struct{}

// Contract of Product Handler
type ProductHandlerInterface interface {
	CreateProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
	GetSpecificProduct(c *gin.Context)
	QueryProducts(c *gin.Context)
	UpdateSpecificProduct(c *gin.Context)
	DeleteSpecificProduct(c *gin.Context)
	DeleteProductsWithIds(c *gin.Context)
}

func GetProductHandler() ProductHandlerInterface {
	if productHandler == nil {
		productHandler = &ProductHandler{}
	}
	return productHandler
}

// Func to Create Product, similar to #Register
func (handler *ProductHandler) CreateProduct(c *gin.Context) {
	var createProductRequest validator.CreateProductRequest
	err := c.ShouldBind(&createProductRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to create new product", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	productRepo := repository.GetProductRepository()
	productModel := &models.Product{}

	// smapping the struct
	smapping.FillStruct(productModel, smapping.MapFields(&createProductRequest))

	if newProduct, err := productRepo.Create(*productModel); err != nil {
		response := response.BuildFailedResponse("failed to create new product", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	} else {
		response := response.BuildSuccessResponse("success to create new product", newProduct)
		c.JSON(http.StatusOK, response)
		return
	}
}

// Func to GetAll Product without in server pagination
func (handler *ProductHandler) GetAllProduct(c *gin.Context) {
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
func (handler *ProductHandler) GetSpecificProduct(c *gin.Context) {
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
func (handler *ProductHandler) QueryProducts(c *gin.Context) {
	pagination := helpers.Pagination{}
	productRepo := repository.GetProductRepository()
	queryPageLimit, isPageLimitExist := c.GetQuery("limit")
	queryPage, isPageQueryExist := c.GetQuery("page")

	if isPageQueryExist {
		pagination.Page, _ = strconv.Atoi(queryPage)
	}

	if isPageLimitExist {
		pagination.Limit, _ = strconv.Atoi(queryPageLimit)
	}

	products, err := productRepo.Query(pagination)

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", products)
	c.JSON(http.StatusOK, response)
}

// Func to Update Product,
func (handler *ProductHandler) UpdateSpecificProduct(c *gin.Context) {
	var updateRequest validator.UpdateProductRequest
	err := c.ShouldBind(&updateRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a product", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	updateModel := &models.Product{}

	// smapping the update request to models
	updateModel.ID = helpers.ConvertStringtoUint(c.Param("productId"))
	smapping.FillStruct(updateModel, smapping.MapFields(&updateRequest))

	productRepo := repository.GetProductRepository()
	err = productRepo.Update(updateModel)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a product", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Specific Product
func (handler *ProductHandler) DeleteSpecificProduct(c *gin.Context) {
	deleteModel := &models.Product{}
	deleteModel.ID = helpers.ConvertStringtoUint(c.Param("productId"))

	productRepo := repository.GetProductRepository()

	err := productRepo.Delete(deleteModel)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete a product", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Products with array ids
func (handler *ProductHandler) DeleteProductsWithIds(c *gin.Context) {
	var deleteRequest validator.DeleteProductsRequest
	err := c.ShouldBind(&deleteRequest)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete products", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	productRepo := repository.GetProductRepository()
	err = productRepo.DeleteWithIds(deleteRequest.Ids)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete products", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
