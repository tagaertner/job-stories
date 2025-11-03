package database

import(
	"fmt"
	"log"
	"os"
	 "github.com/tagaertner/job-stories/services/stories/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect() *gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_PORT"),
	)


	maxRetries := 20
	retryDelay := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Connected to PostgreSQL successfully")
			return db
		}

		log.Printf("❌ Database connection attempt %d/%d failed: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("⏳ Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	log.Fatalf("❌ Could not connect to database after %d retries", maxRetries)
	return nil
}

func RunMigrations(db *gorm.DB){
	db.AutoMigrate(&models.JobStory{},
	&models.UserLog{})
}