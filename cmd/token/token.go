package main

import (
	"fmt"
	"samsamoohooh-api/pkg/config"
	"samsamoohooh-api/pkg/token"
)

func main() {
	cfg, err := config.New("./configs/env.yaml")
	if err != nil {
		panic(err)
	}

	tokenService := token.New(cfg)

	tokenString, err := tokenService.GenerateToken(token.GenerateTokenParams{
		Kind:   token.KindAccess,
		Per:    token.PermissionUser,
		UserID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("generate token:", tokenString)
}
