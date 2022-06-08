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
	userHandler *UserHandler
)

// Func to implement contract of UserHandler
type UserHandler struct{}

// Contract of User Handler
type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	GetSpecificUser(c *gin.Context)
	QueryUsers(c *gin.Context)
	UpdateSpecificUser(c *gin.Context)
	DeleteSpecificUser(c *gin.Context)
	DeleteUsersWithIds(c *gin.Context)
}

func GetUserHandler() UserHandlerInterface {
	if userHandler == nil {
		userHandler = &UserHandler{}
	}
	return userHandler
}

// Func to Create User, similar to #Register
func (handler *UserHandler) CreateUser(c *gin.Context) {
	// var uploadSecureUrl = ""
	var createUserRequest validator.RegisterRequest
	err := c.ShouldBind(&createUserRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to register new user", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userRepo := repository.GetUserRepository()
	// passwordHelper := crypto.GetPasswordCryptoHelper()
	userModel := &models.User{}

	// smapping the struct
	smapping.FillStruct(userModel, smapping.MapFields(&createUserRequest))
	// userModel.Password, _ = passwordHelper.HashAndSalt([]byte(createUserRequest.Password))

	if newUser, err := userRepo.Create(*userModel); err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	} else {
		response := response.BuildSuccessResponse("success login", newUser)
		c.JSON(http.StatusOK, response)
		return
	}
}

// Func to GetAll User without in server pagination
func (handler *UserHandler) GetAllUser(c *gin.Context) {
	userRepo := repository.GetUserRepository()

	users, err := userRepo.GetAll()
	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", users)
	c.JSON(http.StatusOK, response)
}

// Func to GetSpecific User
func (handler *UserHandler) GetSpecificUser(c *gin.Context) {
	userRepo := repository.GetUserRepository()

	user, err := userRepo.GetById(c.Param("userId"))

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", user)
	c.JSON(http.StatusOK, response)
}

// Func to Query User with pagination
func (handler *UserHandler) QueryUsers(c *gin.Context) {
	pagination := helpers.Pagination{}
	userRepo := repository.GetUserRepository()
	queryPageLimit, isPageLimitExist := c.GetQuery("limit")
	queryPage, isPageQueryExist := c.GetQuery("page")

	if isPageQueryExist {
		pagination.Page, _ = strconv.Atoi(queryPage)
	}

	if isPageLimitExist {
		pagination.Limit, _ = strconv.Atoi(queryPageLimit)
	}

	users, err := userRepo.Query(pagination)

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", users)
	c.JSON(http.StatusOK, response)
}

// Func to Update User,
func (handler *UserHandler) UpdateSpecificUser(c *gin.Context) {
	var updateRequest validator.UpdateUserRequest
	err := c.ShouldBind(&updateRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a user", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userRepo := repository.GetUserRepository()

	// smapping the update request to models
	existedUser, err := userRepo.GetById(c.Param("userId"))

	// If user is not exist
	if err != nil {
		response := response.BuildFailedResponse("no such user exist", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
	smapping.FillStruct(existedUser, smapping.MapFields(&updateRequest))

	err = userRepo.Update(existedUser)

	if err != nil {
		response := response.BuildFailedResponse("failed to update an user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Specific User
func (handler *UserHandler) DeleteSpecificUser(c *gin.Context) {
	userRepo := repository.GetUserRepository()
	existedUser, err := userRepo.GetById(c.Param("userId"))

	if err != nil {
		response := response.BuildFailedResponse("no such user exist", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	err = userRepo.Delete(existedUser)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete an user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete User with array ids
func (handler *UserHandler) DeleteUsersWithIds(c *gin.Context) {
	var deleteRequest validator.DeleteUsersRequest
	err := c.ShouldBind(&deleteRequest)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete users", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userRepo := repository.GetUserRepository()
	err = userRepo.DeleteWithIds(deleteRequest.Ids)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete users", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
