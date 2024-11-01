package db

import (
	"fmt"
	"go-api/pkg/shared/config"
	"go-api/pkg/shared/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Exec(`DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'operation_type') THEN
				CREATE TYPE operation_type AS ENUM ('asset', 'liability');
			END IF;
		END $$;`)
	db.AutoMigrate(&models.Account{}, &models.Operation{}, &models.Transaction{})

	return db
}
