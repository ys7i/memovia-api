package infrastructure

import (
	"github.com/ys7i/memovia/interfaces/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func NewDB() *database.DBRepository {
	c := NewConfig()
	return newDB(DBConfig{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func newDB(d DBConfig) *database.DBRepository {
	dsn := "host=db user=" + d.Username + " password=" + d.Password + " dbname=" + d.DBName + " port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	userRepository := &database.UserRepository{DB: db}
	return &database.DBRepository{UserRepo: userRepository}
}
