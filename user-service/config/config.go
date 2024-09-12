package config

import (
    "log"
    "user-service/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

func SetupDatabase() *gorm.DB {
    // Load the database connection from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    sslMode := os.Getenv("SSL_MODE")
	

    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " pgbouncer=true"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.AutoMigrate(&models.User{})
    return db
}
