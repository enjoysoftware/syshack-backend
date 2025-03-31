package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetButterflies(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		googleID := ctx.Param("google_id")

		userID, err := getUserIDByGoogleID(db, googleID)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}

		var butterflies []Butterfly
		if err := db.Where("feed_user_id = ?", userID).Find(&butterflies).Error; err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to get butterflies"})
			return
		}

		ctx.JSON(200, gin.H{
			"butterflies": butterflies,
		})
	}
}

func GetButterfly(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		var butterfly Butterfly
		if err := db.Where("id = ?", id).First(&butterfly).Error; err != nil {
			ctx.JSON(404, gin.H{"error": "Butterfly not found"})
			return
		}

		ctx.JSON(200, butterfly)
	}
}

func CreateButterfly(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var butterfly Butterfly
		google_id := ctx.Param("google_id")
		feed_user_id, err := getUserIDByGoogleID(db, google_id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid feed_user_id or user not found"})
			return
		}

		butterfly.FeedUserID = feed_user_id

		db.Create(&butterfly)

		ctx.JSON(200, butterfly)
	}
}

func updateButterflyState(db *gorm.DB, id uuid.UUID) error {
	var butterfly Butterfly
	if err := db.First(&butterfly, "id = ?", id).Error; err != nil {
		return err
	}

	butterfly.GrowthStage++
	rand.Seed(time.Now().UnixNano())
	butterfly.ColorID += rand.Intn(4)
	fmt.Println(butterfly.ColorID)
	result := db.Model(&butterfly).Updates(map[string]interface{}{"growth_stage": butterfly.GrowthStage, "color_id": butterfly.ColorID})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateButterfly(db *gorm.DB) gin.HandlerFunc { //基本APIで呼ばない
	return func(ctx *gin.Context) {
		butterfly_id := ctx.Param("id")

		id, err := uuid.Parse(butterfly_id)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		err = updateButterflyState(db, id)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to update butterfly growth stage"})
			return
		}

		ctx.JSON(200, gin.H{"message": "OK"})
	}
}
