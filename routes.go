package main

//apiエンドポイントを記載するファイル
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", Hello(db))
	r.GET("/users", GetUsers(db))
	r.GET("/user/:user_id", GetUser(db))
	r.POST("/user", CreateUser(db))
	r.DELETE("/user/:user_id", DeleteUser(db))
	r.PUT("/user/:user_id/administrator", UpdateUserAdministrator(db))

	r.GET("/kakomons", GetKakomons(db))
	r.GET("/kakomon/:id", GetKakomon(db))
	r.POST("/kakomon", CreateKakomon(db))
	r.DELETE("/kakomon/:id", DeleteKakomon(db))

	r.GET("/butterflies/:feed_user_id", GetButterflies(db))
	r.GET("/butterfly/:id", GetButterfly(db))
	r.POST("/butterfly/:id", CreateButterfly(db))
	r.PUT("/butterfly/:id", UpdateButterfly(db))
}
