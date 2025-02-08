package pagination

const (
	DefaultPageSize = 10
	MaxPageSize     = 100
)

type Pagination struct {
	Page int32 `json:"page"`
	Size int32 `json:"size"`
}

func NewPagination(page int32, size int32) Pagination {
	if page < 1 {
		page = 1
	}

	if size < 1 || size > MaxPageSize {
		size = DefaultPageSize
	}

	return Pagination{
		Page: page,
		Size: size,
	}
}

func (pagination *Pagination) Limit() int32 {
	return pagination.Size
}

func (pagination *Pagination) Offset() int32 {
	return (pagination.Page - 1) * pagination.Size
}
