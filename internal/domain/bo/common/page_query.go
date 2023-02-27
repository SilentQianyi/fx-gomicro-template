package boCommon

type PageQuery struct {
	pageNo   int64
	pageSize int64
}

func NewPageQuery(pageNo, pageSize int64) *PageQuery {
	return &PageQuery{
		pageNo:   pageNo,
		pageSize: pageSize,
	}
}

func (p *PageQuery) PageNo() int64 {
	return p.pageNo
}

func (p *PageQuery) PageSize() int64 {
	return p.pageSize
}

func (p *PageQuery) Offset() int64 {
	return (p.pageNo - 1) * p.pageSize
}

func (p *PageQuery) Limit() int64 {
	return p.pageSize
}
