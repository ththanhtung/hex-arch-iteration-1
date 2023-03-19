package main

import (
	"hexArch/internal/adapters/app/api"
	"hexArch/internal/adapters/core/arithmetic"
	"hexArch/internal/adapters/framework/right/db"
	"hexArch/internal/adapters/framework/left/grpc"
	"log"
)

func main() {
	core := arithmetic.NewArithmetic()

	dbAdapter,err := db.NewAdapter("mysql", "admin:passwd")

	if err != nil {
		log.Fatalf("")
	}
	
	appAdapter := api.NewAdapter(dbAdapter, core)

	grpcAdapter := grpc.NewAdapter(appAdapter)

	grpcAdapter.Run()
}