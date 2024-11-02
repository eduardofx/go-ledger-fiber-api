package integration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=ledger_api port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Account{}, &Operation{}, &Transaction{})
	return db, nil
}
