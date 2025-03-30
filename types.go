package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"user_id"`
	Name               string    `json:"name"`
	GoogleID           string    `json:"google_id" gorm:"column:google_id"`
	PreviousUploadDate time.Time `json:"previous_upload_date"`
	IsAdministrator    bool      `json:"is_administrator" gorm:"default:false"`
	CountPost          uint      `json:"count_post" gorm:"default:0"`
	FeedingButterflyID int       `json:"feeding_butterfly_id" gorm:"column:feeding_butterfly_id"`
}

type Kakomon struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Path         string    `json:"path"`
	Grade        string    `json:"grade"`
	Subject      string    `json:"subject"`
	Title        string    `json:"title"` //年_専攻_中身名
	Year         uint      `json:"year"`
	Teacher      string    `json:"teacher"`
	Major        string    `json:"major"`
	UploadUserID uuid.UUID `json:"upload_user_id"`
}

type Butterfly struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	FeedUserID  uuid.UUID `gorm:"column:feed_user_id" json:"feed_user_id"`
	GrowthStage int       `json:"growth_stage"`
	ColorID     int       `json:"color_id"`
	UpdateDate  time.Time `gorm:"column:update_date" json:"update_date"`
}
