package infrastructure

import (
	"fmt"

	"github.com/ys7i/memorizer/src/interfaces/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host string;
	Username string
	Password string
	DBName string
}

func NewDB() *database.DBRepository {
	c := NewConfig()
	return newDB(DBConfig{
		Host: c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName: c.DB.Production.DBName,
	})
}

func newDB(d DBConfig) *database.DBRepository {
	dsn := "tcp://db:5432?database=" + d.DBName + "&username=" + d.Username + "&password=" + d.Password + "&read_timeout=10&write_timeout=2"
	fmt.Printf(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	userRepository := &database.UserRepository{DB: db}
	return &database.DBRepository{UserRepo: userRepository}
}