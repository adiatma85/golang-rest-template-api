// Code generated by mockery v2.10.0. DO NOT EDIT.

package handler

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// AuthHandlerInterface is an autogenerated mock type for the AuthHandlerInterface type
type AuthHandlerInterface struct {
	mock.Mock
}

// AuthLogin provides a mock function with given fields: c
func (_m *AuthHandlerInterface) AuthLogin(c *gin.Context) {
	_m.Called(c)
}

// AuthRegister provides a mock function with given fields: c
func (_m *AuthHandlerInterface) AuthRegister(c *gin.Context) {
	_m.Called(c)
}
