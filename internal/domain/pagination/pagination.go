package pagination

import (
	"fmt"
	"strings"
)

const (
	MaxPageSize     = 100
	DefaultPageSize = 10
)

type Pagination struct {
	OrderBy  string         `json:"orderBy"`
	OrderDir OrderDirection `json:"orderDir"`
	Page     int32          `json:"page"`
	Size     int32          `json:"size"`
}

func NewPagination(page int32, size int32, orderBy string, orderDir OrderDirection) Pagination {
	orderDir = OrderDirection(strings.ToUpper(orderDir.String()))
	if orderDir != ASC && orderDir != DESC {
		orderDir = ASC
	}

	if page < 1 {
		page = 1
	}

	if size < 1 || size > MaxPageSize {
		size = DefaultPageSize
	}

	return Pagination{
		OrderBy:  orderBy,
		OrderDir: orderDir,
		Page:     page,
		Size:     size,
	}
}

func (pagination *Pagination) Order() string {
	if pagination.OrderBy == "" {
		return ""
	}
	return fmt.Sprintf("%s %s", pagination.OrderBy, pagination.OrderDir)
}

func (pagination *Pagination) Limit() int32 {
	return pagination.Size
}

func (pagination *Pagination) Offset() int32 {
	return (pagination.Page - 1) * pagination.Size
}
