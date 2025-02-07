package pagination

type OrderDirection string

const (
	ASC  OrderDirection = "ASC"
	DESC OrderDirection = "DESC"
)

func (orderDirection OrderDirection) String() string {
	return string(orderDirection)
}
