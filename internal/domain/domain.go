package domain

type Pagination struct {
	Page    *int `form:"page,default=1" json:"page,omitempty" binding:"gte=1"`
	PerPage *int `form:"perPage,default=10" json:"per_page,omitempty" binding:"gte=1,lte=1000"`
}
