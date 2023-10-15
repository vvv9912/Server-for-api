package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"serverBot2/internal/server"
	"serverBot2/internal/storage"
)

func main() {
	//a := []byte("aaa")
	//b := []byte("ccc")
	//c := make([][]byte, 0)
	//c = append(c, a)
	//c = append(c, b)

	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(db)
		return
	}
	sProducts := storage.NewProductsPostgresStorage(db)

	s := server.NewServer(sProducts)
	ctx := context.Background()
	err = s.ServerStart(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
