// Code generated by mockery v2.10.0. DO NOT EDIT.

package helpers

import mock "github.com/stretchr/testify/mock"

// PasswordCryptoHelper is an autogenerated mock type for the PasswordCryptoHelper type
type PasswordCryptoHelper struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashPassword, plainPassword
func (_m *PasswordCryptoHelper) ComparePassword(hashPassword string, plainPassword []byte) bool {
	ret := _m.Called(hashPassword, plainPassword)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []byte) bool); ok {
		r0 = rf(hashPassword, plainPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HashAndSalt provides a mock function with given fields: pwd
func (_m *PasswordCryptoHelper) HashAndSalt(pwd []byte) (string, error) {
	ret := _m.Called(pwd)

	var r0 string
	if rf, ok := ret.Get(0).(func([]byte) string); ok {
		r0 = rf(pwd)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
