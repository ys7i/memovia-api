package main

import (
	"log"

	"github.com/ys7i/memorizer/src/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	// db.UserRepo.DB.AutoMigrate(&domain.User{})
	log.Print("hello")
	r := infrastructure.NewRouting(db)
  r.Run()
}