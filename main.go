package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"serverBot2/internal/config"
	"serverBot2/internal/server"
	"serverBot2/internal/storage"
)

func main() {

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	defer db.Close()

	if err != nil {
		fmt.Println(db)
		return
	}

	sProducts := storage.NewProductsPostgresStorage(db)

	s := server.NewServer(sProducts, sProducts)
	ctx := context.Background()

	serverAddr := config.Get().ServerAddress
	err = s.ServerStart(ctx, serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
}
