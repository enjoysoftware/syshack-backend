package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	user_id              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	name                 string
	google_id            string
	previous_upload_date time.Time
	is_administrator     bool
	feeding_butterfly_id int
}

type Kakomon struct {
	gorm.Model
	id             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	path           string
	grade          string
	subject        string
	year           int
	major          string
	title          string
	upload_user_id uuid.UUID
}

type Butterfly struct {
	gorm.Model
	id           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	feed_user_id uuid.UUID
	growth_stage int
	color_id     int
	update_date  time.Time
}
