package utils

import (
	"yuedan/enums"
)

type DataWrapper struct {
	Status     int
	Error      enums.Errors
	Data       interface{}
	PageSize   int
	PageIndex  int
	TotalPage  int
	TotalCount int
}

func GetTotalPage(totalCount int, pageSize int) int {
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
func GenerateDataWrapperWithPage(errorMessage enums.Errors, data interface{}, pageSize int, pageIndex int, totalPage int, totalCount int) DataWrapper {
	var status int
	if errorMessage != "" {
		status = 1
	}
	return DataWrapper{Status: status, Error: errorMessage, Data: data, PageSize: pageSize, PageIndex: pageIndex, TotalPage: totalPage, TotalCount: totalCount}
}
