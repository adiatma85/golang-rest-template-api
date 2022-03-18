package handler

import (
	"net/http"
	"strconv"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/repository"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/validator"
	"github.com/adiatma85/go-tutorial-gorm/pkg/crypto"
	"github.com/adiatma85/go-tutorial-gorm/pkg/helpers"
	"github.com/adiatma85/go-tutorial-gorm/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

// Func to Create User, similar to #Register
func CreateUser(c *gin.Context) {
	var createUserRequest validator.RegisterRequest
	err := c.ShouldBind(&createUserRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to register new user", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userRepo := repository.GetUserRepository()
	passwordHelper := crypto.GetPasswordCryptoHelper()
	userModel := &models.User{}

	// smapping the struct
	smapping.FillStruct(userModel, smapping.MapFields(&createUserRequest))
	userModel.Password, _ = passwordHelper.HashAndSalt([]byte(createUserRequest.Password))

	if newUser, err := userRepo.Create(*userModel); err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		response := response.BuildSuccessResponse("success login", newUser)
		c.JSON(http.StatusOK, response)
		return
	}
}

// Func to GetAll User without in server pagination
func GetAllUser(c *gin.Context) {
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
func GetSpecificUser(c *gin.Context) {
	userRepo := repository.GetUserRepository()

	user, err := userRepo.GetById(c.Param("userId"))

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", user)
	c.JSON(http.StatusOK, response)
}

// Func to Query User with pagination
func QueryUsers(c *gin.Context) {
	pagination := helpers.Pagination{}
	userRepo := repository.GetUserRepository()
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

	users, err := userRepo.Query(&models.User{}, pagination)

	if err != nil {
		response := response.BuildFailedResponse("failed to fetch data", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := response.BuildSuccessResponse("success to fetch data", users)
	c.JSON(http.StatusOK, response)
}

// Func to Update User,
func UpdateSpecificUser(c *gin.Context) {
	var updateRequest validator.UpdateUserRequest
	err := c.ShouldBind(&updateRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to update a user", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	updateModel := &models.User{}

	// smapping the update request to models
	updateModel.ID, _ = strconv.ParseUint(c.Param("userId"), 10, 64)
	smapping.FillStruct(updateModel, smapping.MapFields(&updateRequest))

	userRepo := repository.GetUserRepository()
	err = userRepo.Update(updateModel)

	if err != nil {
		response := response.BuildFailedResponse("failed to update an user", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete Specific User
func DeleteSpecificUser(c *gin.Context) {
	deleteModel := &models.User{}
	deleteModel.ID, _ = strconv.ParseUint(c.Param("userId"), 10, 64)

	userRepo := repository.GetUserRepository()

	err := userRepo.Delete(deleteModel)
	if err != nil {
		response := response.BuildFailedResponse("failed to delete an user", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Func to Delete User with array ids
func DeleteUsersWithIds(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "ok",
		"message": "need revision for delete users",
	})
}
