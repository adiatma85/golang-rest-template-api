package repository

import "github.com/adiatma85/golang-rest-template-api/pkg/helpers"

// variation of pagination type
var (
	// default pagination
	pagination0 helpers.Pagination = helpers.Pagination{}

	// With limit and page
	pagination1 helpers.Pagination = helpers.Pagination{
		Limit: 5,
		Page:  1,
	}

	// Another sort like (asc or desc)
	pagination2 helpers.Pagination = helpers.Pagination{
		Sort: "",
	}
)
