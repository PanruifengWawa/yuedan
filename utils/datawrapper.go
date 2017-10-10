package utils

import (
	"yuedan/enums"
)

type DataWrapper struct {
	Status     int
	Error      enums.Errors
	Data       interface{}
	PageSize   int64
	PageIndex  int64
	TotalPage  int64
	TotalCount int64
}

func GetTotalPage(totalCount int64, pageSize int64) int64 {
	if totalCount%pageSize == 0 {
		return totalCount / pageSize
	} else {
		return totalCount/pageSize + 1
	}
}
func GenerateDataWrapper(errorMessage enums.Errors, data interface{}) DataWrapper {
	var status int
	if errorMessage != "" {
		status = 1
	}
	return DataWrapper{Status: status, Error: errorMessage, Data: data}
}
func GenerateDataWrapperWithPage(errorMessage enums.Errors, data interface{}, pageSize int64, pageIndex int64, totalCount int64) DataWrapper {
	var status int
	if errorMessage != "" {
		status = 1
	}
	totalPage := GetTotalPage(totalCount, pageSize)
	return DataWrapper{Status: status, Error: errorMessage, Data: data, PageSize: pageSize, PageIndex: pageIndex, TotalPage: totalPage, TotalCount: totalCount}
}
