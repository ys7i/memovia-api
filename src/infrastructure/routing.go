package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/ys7i/memovia/interfaces/controllers"
	"github.com/ys7i/memovia/interfaces/database"
)

type Routing struct {
	DB   *database.DBRepository
	Gin  *gin.Engine
	Port string
}

func NewRouting(dbRepo *database.DBRepository) *Routing {
	r := &Routing{
		DB:   dbRepo,
		Gin:  gin.Default(),
		Port: ":5050",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
	r.Gin.POST("/users/create", func(c *gin.Context) { usersController.Create(c) })
	r.Gin.POST("users/apple-signin", func(c *gin.Context) { usersController.GetAppleUser(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
