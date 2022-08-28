package handler

import (
	"gRPC_Users/internal/service"
	"gRPC_Users/pkg"
)

type Handler struct {
	pkg.UnimplementedUserApiServer
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
