package infrastructure

import (
	"os"
)

type Config struct {
    DB struct {
        Production struct {
            Host string
            Username string
            Password string
            DBName string
        }
        Test struct {
            Host string
            Username string
            Password string
            DBName string
        }
    }
    Routing struct {
        Port string
    }
}

func NewConfig() *Config {
    c := new(Config)
    c.DB.Production.Host = "db"
    c.DB.Production.Username = os.Getenv("POSTGRES_USER")
    c.DB.Production.Password = os.Getenv("POSTGRES_PASSWORD")
    c.DB.Production.DBName = os.Getenv("POSTGRES_DB")

    c.DB.Test.Host = "db"
    c.DB.Test.Username = "test"
    c.DB.Test.Password = "password"
    c.DB.Test.DBName = "db_test"
    
    c.Routing.Port = ":5432"

    return c
}