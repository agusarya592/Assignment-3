package utils

import (
	"assignment3/utils/dto"
	"assignment3/utils/errors"
	"encoding/json"
	"log"
	"net/http"
)

func NewSuccessResponsWriter(rw http.ResponseWriter, code int, status string, data interface{}) {
	BaseResponseWriter(rw, code, status, nil, data)
}

func NewErrorResponse(rw http.ResponseWriter, err error) {
	errMap := errors.GetErrorResponseMetaData(err)
	BaseResponseWriter(rw, errMap.Code, "", &dto.ErrorData{Code: errMap.Code, Message: errMap.Message}, nil)
}

func BaseResponseWriter(rw http.ResponseWriter, code int, status string, er *dto.ErrorData, data interface{}) {
	res := dto.BaseResponse{
		Status: status,
		Data:   data,
		Error:  er,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Printf("cant marshal the interface")
	}
	rw.Header().Add("Content Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(jsonData)
}
