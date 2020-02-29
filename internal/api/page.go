package api

// Page - the common API entity for pagination cases.
type Page struct {
	CurrentPage   int           `json:"current_page"`
	PageSize      int           `json:"page_size"`
	TotalPages    int           `json:"total_pages"`
	TotalElements int64         `json:"total_elements"`
	Elements      []interface{} `json:"elements"`
}
