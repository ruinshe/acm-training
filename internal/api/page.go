package api

// RequestPage - the page informatin passed from UI side.
type RequestPage struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// Page - the common API entity for pagination cases.
type Page struct {
	CurrentPage   int         `json:"current_page"`
	PageSize      int         `json:"page_size"`
	TotalPages    int         `json:"total_pages"`
	TotalElements int         `json:"total_elements"`
	Elements      interface{} `json:"elements"`
}
