package selector

const (
	defaultLimit = 20
)

// Paging Model
type Paging struct {
	Skip      int64
	Limit     int64
	SortKey   string
	SortVal   int
	Condition interface{}
}

func NewPaging() *Paging {
	return &Paging{
		Limit: defaultLimit,
	}
}

func (p *Paging) SetPaging(page, limit int64) *Paging {
	if limit > 0 {
		p.Limit = limit
	}
	if page > 0 {
		p.Skip = (page - 1) * p.Limit
	}
	return p
}

func (p *Paging) SetSorting(key string, val int) *Paging {
	p.SortKey = key
	p.SortVal = val
	return p
}
