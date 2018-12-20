package models

type Page struct {
	pageSize int
	pageNo int
	total int
	rows interface{}
}

func (page *Page) SetPageSize(pageSize int)  {
	page.pageSize = pageSize
}

func (page *Page) SetPageNo(pageNo int)  {
	page.pageNo = pageNo
}

func (page *Page) GetTotal() int {
	return page.total
}

func (page *Page) SetTotal(total int)  {
	page.total = total
}

func (page *Page) GetRows() interface{} {
	return page.rows
}

func (page *Page) SetRows(rows interface{})  {
	page.rows = rows
}

func (page *Page) GetOffset() int {
	return (page.pageNo-1) * page.pageSize;
}
