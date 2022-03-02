package body

import (
	"time"

	"github.com/uccu/go-stringify"
)

type DataBase struct {
	ErrorCode int    `json:"code" example:"0"`
	ErrorMsg  string `json:"msg" example:"ok"`
}

type InfoData struct {
	DataBase
	Data map[string]interface{} `json:"data"`
}

type ListData struct {
	DataBase
	Data struct {
		List  []map[string]interface{} `json:"list"`
		Count int64                    `json:"count"`
	} `json:"data"`
}

type PageTrait struct {
	Page   uint `json:"page" uri:"page" example:"1"`     // 页数
	Length uint `json:"length" uri:"length" example:"1"` // 每页数量
}

type Page interface {
	GetOffset() uint
	GetPage() uint
	GetLength() uint
}

type IdParam struct {
	Id int64 `json:"id" binding:"required"`
}

func (p PageTrait) GetOffset() uint {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Length * (p.Page - 1)
}

func (p PageTrait) GetPage() uint {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p PageTrait) GetLength() uint {
	return p.Length
}

type Date interface {
	GetStart() *time.Time
	GetEnd() *time.Time
}

type DateTrait struct {
	StartDate *string `json:"startDate"` // 开始日期, 格式: YYYY-MM-DD
	EndDate   *string `json:"endDate"`   // 结束日期, 格式: YYYY-MM-DD
}

func (p DateTrait) GetStart() *time.Time {
	return parseDate(p.StartDate)
}

func (p DateTrait) GetEnd() *time.Time {
	return parseDate(p.EndDate)
}

type DateTime interface {
	GetStart() *time.Time
	GetEnd() *time.Time
}

func parseDate(i *string) *time.Time {
	if i == nil {
		return nil
	}

	sli := stringify.ToIntSlice(*i, "-")
	if len(sli) != 3 {
		return nil
	}

	date := time.Date(int(sli[0]), time.Month(sli[1]), int(sli[2]), 0, 0, 0, 0, time.Local)
	return &date
}

type DateTimeTrait struct {
	StartTime *string `json:"startTime"` // 开始日期, 格式: YYYY-MM-DD HH:ii:ss
	EndTime   *string `json:"endTime"`   // 结束日期, 格式: YYYY-MM-DD HH:ii:ss
}

func (p DateTimeTrait) GetStart() *time.Time {
	return parseTime(p.StartTime)
}

func (p DateTimeTrait) GetEnd() *time.Time {
	return parseTime(p.EndTime)
}

func parseTime(i *string) *time.Time {
	if i == nil || *i == "" {
		return nil
	}

	sli := stringify.ToStringSlice(*i, " ")

	day := stringify.ToIntSlice(sli[0], "-")
	if len(day) != 3 {
		return nil
	}

	var hour int
	var min int
	var sec int

	if len(sli) > 1 {
		ti := stringify.ToIntSlice(sli[1], ":")
		if len(ti) > 0 {
			hour = int(ti[0])
		}
		if len(ti) > 1 {
			min = int(ti[1])
		}
		if len(ti) > 2 {
			sec = int(ti[2])
		}
	}

	date := time.Date(int(day[0]), time.Month(day[1]), int(day[2]), hour, min, sec, 0, time.Local)
	return &date
}
