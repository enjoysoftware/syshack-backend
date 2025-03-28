package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("データベース接続エラー: %w", err)
	}
	return db, nil
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Kakomon{}, &Butterfly{})
	if err != nil {
		return fmt.Errorf("データベースのマイグレーションエラー: %w", err)
	}
	return nil
}

func main() {
	r := gin.Default()

	// データベースに接続
	db, err := SetupDatabase()
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}

	// データベースにテーブルを作成
	err = MigrateDatabase(db)
	if err != nil {
		log.Fatal("データベースのマイグレーションエラー:", err)
	}

	// ルーティング設定
	SetupRoutes(r, db)

	r.Run(":8080")
}
