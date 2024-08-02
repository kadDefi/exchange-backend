package domain

type Pagination struct {
	Page    *int `form:"page,default=0" json:"page,omitempty" binding:"gte=0"`
	PerPage *int `form:"perPage,default=10" json:"per_page,omitempty" binding:"gte=1,lte=1000"`
}
