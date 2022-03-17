package helpers

// Pagination struct for query result
// For more reference, you can access https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f
type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	TotalRows  int64       `json:"total_rows,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Rows       interface{} `json:"rows"`
}

// Func to get Offset for querying
func (p *Pagination) GetOffset() int {
	return 0
}

// Func to get Limit for querying. Default limit is 10
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

// Func to get Offset for querying
func (p *Pagination) GetPage() int {
	return 0
}
