package databases

import (
	"fmt"
	"os"
	"testing"

	"github.com/abelherl/go-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func ConnectToTestDB(t *testing.T) *gorm.DB {
	dsn := os.Getenv("DB_STRING")
	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("failed to migrate test schema: %v", err)
	}

	TestDB = db
	return db
}
