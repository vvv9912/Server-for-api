package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"serverBot2/internal/server"
	"serverBot2/internal/storage"
)

type FakeDb struct {
}

func (f FakeDb) Connect() *sqlx.DB {
	return nil
}
func main() {

	//db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	//defer db.Close()

	//if err != nil {
	//	fmt.Println(db)
	//	return
	//}
	fake := FakeDb{}
	db := fake.Connect()
	sProducts := storage.NewProductsPostgresStorage(db)

	s := server.NewServer(sProducts)
	ctx := context.Background()
	err := s.ServerStart(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}
