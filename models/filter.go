package models

import "math"

type Filter struct {
	Page, PageSize int
}

type SearchParm struct {
	Product   string
	Author    string
	Categorty string
	OrderBY   string
	Oprice    string
}

func (f Filter) Limit() int {
	return f.PageSize
}

func (f Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

type Metadata struct {
	CurrentPage, PageSize, FirstPage, LastPage, TotalRecords int
}

func ComputeMetadata(toatalRecords, page, pageSize int) Metadata {
	if toatalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(toatalRecords) / float64(pageSize))),
		TotalRecords: toatalRecords,
	}
}
