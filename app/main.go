package main

import (
	"context"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/inmemorydb"
	"log"
)

func main() {
	ctx := context.Background()
	service := initializeMockService()

	user, err := service.Create(ctx, "john")
	if err != nil {
		log.Fatal(err)
	}

	_, err = service.Show(ctx, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	inmemorydb.Dump()
}
