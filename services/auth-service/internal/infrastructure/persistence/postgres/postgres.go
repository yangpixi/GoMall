package postgres

import (
	"fmt"

	"github.com/yangpixi/GoMall/auth-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.DataBase.Connection.Host,
		c.DataBase.Connection.Port,
		c.DataBase.User,
		c.DataBase.Password,
		c.DataBase.DB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return db, err
}
