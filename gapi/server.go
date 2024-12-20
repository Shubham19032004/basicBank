package gapi

import (
	db "bank/db/sqlc"
	"bank/pb"
	"bank/token"
	"bank/utils"
	"fmt"
)

// server serves HTTP requests for our banking service
type Server struct {
    pb.UnimplementedSimpleBankServer
    store      db.Store
    tokenMaker token.Maker
    config     utils.Config
}

// NewServer creates a new HTTP server and setup routing

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker:%w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
