package main

import (
	"github.com/ys7i/memorizer/domain"
	"github.com/ys7i/memorizer/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	db.UserRepo.DB.AutoMigrate(&domain.User{})
	r := infrastructure.NewRouting(db)
  r.Run()
}