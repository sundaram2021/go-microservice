// config/config.go
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// In your SetupDatabase function
func SetupDatabase() (*gorm.DB, error) {
	dsn := "user=postgres.bibabfrfswkkgzrepncb password=xsuu66hwXVQ45Wce host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logs
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
