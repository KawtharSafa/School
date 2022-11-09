package model

type PaginationFilter struct {
	SortBy     string `query:"sortBy"`
	Direction  string `query:"direction"`
	Classroom  string `query:"classroom"`
	Student    string `query:"student"`
	Teacher    string `query:"teacher"`
	TopStudent bool   `query:"topStudent"`
	Page       int    `query:"page"`
	PerPage    bool   `query:"perPage"`
}
