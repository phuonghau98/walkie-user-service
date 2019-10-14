package main

import (
	"context"
	"fmt"
	userPB "github.com/phuonghau98/walkie-user-service/proto"
)

type Handler struct {
}

func (h *Handler) Authenticate(ctx context.Context, req *userPB.User, res *userPB.TokenInfo) error {
	fmt.Println(req)
	return nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *userPB.TokenInfo, res *userPB.TokenInfo) error {
	fmt.Println(req)
	return nil
}


