package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewErrorResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
