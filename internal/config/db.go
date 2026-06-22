package config

import (
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *Config) *gorm.DB{
	dsn := cfg.Dsn 

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("failed to Connect Database")
	}else {
		println("Database Connected Successfully")
	}
	return db 
}