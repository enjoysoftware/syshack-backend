package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getUserIDByGoogleID(db *gorm.DB, googleID string) (uuid.UUID, error) {
	var user User
	if err := db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
		return uuid.Nil, err
	}
	return user.UserID, nil
}
